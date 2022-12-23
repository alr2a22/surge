package ride_req_handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"surge/api/ride/models"
	"surge/internal/db"
	"surge/internal/geo_query"
	"surge/internal/http_response"
	"surge/internal/middlewares"
	"surge/internal/validation"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RideRequestView struct {
	Lat  float32 `json:"latitude" validate:"required"`
	Long float32 `json:"longitude" validate:"required"`
}

type RideResponseView struct {
	Rate float32 `json:"rate"`
}

var ErrNotSupport = errors.New("not supported area")
var ErrServiceWithoutThresholds = errors.New("service without threshols")

func RideHandler(w http.ResponseWriter, r *http.Request) {
	var rrv RideRequestView
	err := json.NewDecoder(r.Body).Decode(&rrv)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	errMap, isValid := validation.Validate(rrv)
	if !isValid {
		http_response.ResponseWithMultipleError(w, http.StatusBadRequest, errMap)
		return
	}

	districtID, err := geo_query.FindDistrict(rrv.Lat, rrv.Long)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusNotAcceptable, err)
		return
	}

	rr := &models.RideRequest{
		Lat:      rrv.Lat,
		Long:     rrv.Long,
		District: districtID,
		UserID:   middlewares.GetUser(r).ID,
	}
	err = rr.Create()
	if err != nil {
		http_response.ResponseWithError(w, http.StatusNotAcceptable, err)
		return
	}

	nReq := db.GetRedisBackend().AddNewRequestWithPrefix(districtID)
	// nReq, _ := rr.CountRideRequestWithDistrictWithWindow(districtID, time.Minute * time.Duration(config.GetConfig().WindowMinutes))

	logrus.Debugln("number requets in window:", nReq)
	coefficient, err := models.GetCurrentCoefficient(nReq)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http_response.ResponseWithError(w, http.StatusNotAcceptable, ErrServiceWithoutThresholds)
			return
		}
		http_response.ResponseWithError(w, http.StatusNotAcceptable, err)
		return
	}
	logrus.Debugln("coefficient is:", coefficient)
	rate := coefficient * 1000

	http_response.ResponseSuccess(w, http.StatusOK, RideResponseView{rate})
}
