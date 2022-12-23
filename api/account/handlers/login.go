package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"surge/api/account/models"
	"surge/internal/db"
	"surge/internal/http_response"
	"surge/internal/middlewares"
	"surge/internal/validation"
)

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

var AuthenticationErr = errors.New("authentication failure")

func Login(w http.ResponseWriter, r *http.Request) {
	DB := db.GetDBConn()

	var input LoginReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	errMap, isValid := validation.Validate(input)
	if !isValid {
		http_response.ResponseWithMultipleError(w, http.StatusBadRequest, errMap)
		return
	}

	var user models.User
	err = DB.First(&user, "username = ?", input.Username).Error
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, AuthenticationErr)
		return
	}

	err = user.CheckPassword(input.Password)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, AuthenticationErr)
		return
	}

	token, err := middlewares.GenerateJWT(user)
	if err != nil {
		http_response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	http_response.ResponseSuccess(w, http.StatusOK, LoginResp{token})
}
