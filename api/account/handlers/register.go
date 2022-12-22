package handlers

import (
	"encoding/json"
	"net/http"
	"surge/api/account/models"
	"surge/internal/db"
	"surge/internal/http_response"
	"surge/internal/validation"
)

func Register(w http.ResponseWriter, r *http.Request) {
	DB := db.GetDBConn()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	errMap, isValid := validation.Validate(user)
	if !isValid {
		http_response.ResponseWithMultipleError(w, http.StatusBadRequest, errMap)
		return
	}

	user.SetPassword(user.Password)

	err = DB.Create(&user).Error
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	http_response.ResponseSuccess(w, http.StatusCreated, user)
}
