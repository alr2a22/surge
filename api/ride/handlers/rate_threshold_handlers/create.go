package rate_threshold_handlers

import (
	"encoding/json"
	"net/http"
	"surge/api/ride/models"
	"surge/internal/http_response"
	"surge/internal/validation"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var rrt models.RideRateThreshold
	err := json.NewDecoder(r.Body).Decode(&rrt)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	errMap, isValid := validation.Validate(rrt)
	if !isValid {
		http_response.ResponseWithMultipleError(w, http.StatusBadRequest, errMap)
		return
	}

	err = rrt.Create()
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	http_response.ResponseSuccess(w, http.StatusCreated, rrt)
}
