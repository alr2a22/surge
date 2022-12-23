package rate_threshold_handlers

import (
	"net/http"
	"surge/api/ride/models"
	"surge/internal/http_response"
)

func List(w http.ResponseWriter, r *http.Request) {
	var rrts models.RideRateThresholdList
	rrts.List()
	http_response.ResponseSuccess(w, http.StatusOK, rrts)
}
