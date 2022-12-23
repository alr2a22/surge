package middlewares

import (
	"errors"
	"net/http"
	"surge/api/account/models"
	"surge/internal/http_response"

	"github.com/gorilla/context"
	"github.com/sirupsen/logrus"
)

var AdministrationErr = errors.New("admin user required")

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uc, ok := context.Get(r, "user").(models.User); !ok || !uc.Admin {
			logrus.Debug("admin checker", ok, uc)
			http_response.ResponseWithError(w, http.StatusUnauthorized, AdministrationErr)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUser(r *http.Request) models.User {
	return context.Get(r, "user").(models.User)
}
