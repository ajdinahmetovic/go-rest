package routes

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/dgrijalva/jwt-go"
)

//VerifyTokenMiddleware func
func VerifyTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.EnableCors(&w, r)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				return []byte("tajna"), nil
			})
			if err != nil {
				v, _ := err.(*jwt.ValidationError)
				if v.Errors == jwt.ValidationErrorSignatureInvalid {
					httputil.WriteError(w, err, http.StatusUnauthorized)
					return
				}
				if v.Errors == jwt.ValidationErrorExpired {
					httputil.WriteError(w, err, http.StatusUnauthorized)
					return
				}
			}
			if token.Valid {
				next(w, r)
			} else {
				httputil.WriteError(w, err, http.StatusUnauthorized)
				return
			}
		} else {
			httputil.WriteError(w, errors.New("Bearer token is required"), http.StatusUnauthorized)
		}
	})
}
