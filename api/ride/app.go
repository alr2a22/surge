package ride

import (
	"surge/api/ride/handlers/rate_threshold_handlers"
	"surge/api/ride/handlers/ride_req_handlers"
	"surge/api/ride/models"
	"surge/internal/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AddRoutes(r *mux.Router) {
	rideSubRouter := r.PathPrefix("/rides").Subrouter()
	rideSubRouter.Use(middlewares.JWTMiddleware)
	rideSubRouter.HandleFunc("", ride_req_handlers.RideHandler).Methods("POST")

	thresholdSubRouter := r.PathPrefix("/thresholds").Subrouter()
	thresholdSubRouter.Use(middlewares.JWTMiddleware)
	thresholdSubRouter.Use(middlewares.AdminMiddleware)
	thresholdSubRouter.HandleFunc("", rate_threshold_handlers.List).Methods("GET")
	thresholdSubRouter.HandleFunc("/{id}", rate_threshold_handlers.Retrieve).Methods("GET")
	thresholdSubRouter.HandleFunc("", rate_threshold_handlers.Create).Methods("POST")
	thresholdSubRouter.HandleFunc("/{id}", rate_threshold_handlers.Update).Methods("PUT")
	thresholdSubRouter.HandleFunc("/{id}", rate_threshold_handlers.Delete).Methods("DELETE")
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(models.RideRateThreshold{})
	DB.AutoMigrate(models.RideRequest{})
}
