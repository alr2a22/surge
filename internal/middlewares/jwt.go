package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"surge/api/account/models"
	"surge/internal/config"
	"surge/internal/db"
	"surge/internal/http_response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/sirupsen/logrus"
)

var AuthenticationRequiredErr = errors.New("authentication required")

type UserClaims struct {
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Admin     bool      `json:"super_user"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) {
	signingKey := []byte(config.GetConfig().JwtSecret)
	t := jwt.New(jwt.SigningMethodHS256)
	d := time.Duration(config.GetConfig().JwtValidDays) * 24 * time.Hour
	t.Claims = UserClaims{
		UserID:    user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
		},
	}

	tokenString, err := t.SignedString(signingKey)
	return tokenString, err
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		signingKey := []byte(config.GetConfig().JwtSecret)
		if len(tokenString) == 0 {
			http_response.ResponseWithError(w, http.StatusForbidden, AuthenticationRequiredErr)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})
		if err != nil {
			logrus.Debugln("token not valid because", err.Error())
			http_response.ResponseWithError(w, http.StatusUnauthorized, AuthenticationRequiredErr)
			return
		}

		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			DB := db.GetDBConn()
			var user models.User
			err = DB.First(&user, "id = ?", claims.UserID).Error
			if err != nil || user.CreatedAt != claims.CreatedAt {
				logrus.Debugln("token not valid because invalid user")
				http_response.ResponseWithError(w, http.StatusUnauthorized, AuthenticationRequiredErr)
				return
			}

			logrus.Debugf("id, username, admin:%v, %v, exire at: %v", claims.UserID, claims.Username, claims.ExpiresAt)
			context.Set(r, "user", user)
		} else {
			logrus.Debugln("token not valid because", err.Error())
			http_response.ResponseWithError(w, http.StatusUnauthorized, AuthenticationRequiredErr)
			return
		}
		next.ServeHTTP(w, r)
	})
}
