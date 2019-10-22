package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"app/pkg/responses"
)

var signingKey = []byte("aqOeh4ck3R")

//Authorized - Middleware to check validation of JWT
func Authorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {

			rawToken := r.Header["Authorization"][0]
			splitToken := strings.Split(rawToken, "Bearer")
			validToken := strings.TrimSpace(splitToken[1])

			token, err := jwt.Parse(validToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			if err != nil {
				responses.Error(w, 400, err.Error())
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			responses.Error(w, 400, "Not Authorized")
		}
	})
}
