package account

import (
	"surge/api/account/handlers"
	"surge/api/account/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(models.User{})
}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
}
