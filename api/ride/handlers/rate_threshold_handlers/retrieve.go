package rate_threshold_handlers

import (
	"net/http"
	"strconv"
	"surge/api/ride/models"
	"surge/internal/http_response"

	"github.com/gorilla/mux"
)

func Retrieve(w http.ResponseWriter, r *http.Request) {
	var rrt models.RideRateThreshold

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	err = rrt.GetByID(id)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusNotFound, err)
		return
	}

	http_response.ResponseSuccess(w, http.StatusOK, rrt)
}
