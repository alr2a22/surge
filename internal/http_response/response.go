package http_response

import (
	"encoding/json"
	"net/http"
)

type ResponseErr struct {
	Error string `json:"error"`
}

type ResponseErrs struct {
	Errors map[string]string `json:"errors"`
}

func ResponseWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponseErr{Error: err.Error()})
}

func ResponseWithMultipleError(w http.ResponseWriter, statusCode int, errMap map[string]string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponseErrs{Errors: errMap})
}

func ResponseSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
