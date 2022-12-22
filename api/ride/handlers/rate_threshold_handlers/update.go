package rate_threshold_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"surge/api/ride/models"
	"surge/internal/http_response"
	"surge/internal/validation"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var rrt models.RideRateThreshold

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	errMap, isValid := validation.Validate(rrt)
	if !isValid {
		http_response.ResponseWithMultipleError(w, http.StatusBadRequest, errMap)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&rrt)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		http_response.ResponseWithError(w, http.StatusNotFound, err)
		return
	}
	rrt.ID = uint(id)

	err = rrt.Save()
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	http_response.ResponseSuccess(w, http.StatusOK, rrt)
}
