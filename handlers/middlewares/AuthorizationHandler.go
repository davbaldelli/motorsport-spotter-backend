package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"motorsportspotter.backend/handlers"
	"motorsportspotter.backend/models"
	"net/http"
	"strings"
)

type AuthorizationHandlerImpl struct {
	Secret string
}

func (a AuthorizationHandlerImpl) Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			r.Header.Set("Role", string(models.Base))
			next.ServeHTTP(w, r)
			return
		}
		var mySigningKey = []byte(a.Secret)

		tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			handlers.RespondError(w, http.StatusUnauthorized, fmt.Errorf("your Token has expired: %v", err.Error()))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			switch claims["role"] {
			case "admin":
				{
					r.Header.Set("Role", string(models.Admin))
					next.ServeHTTP(w, r)
					return
				}
			default:
				{
					r.Header.Set("Role", string(models.Base))
					next.ServeHTTP(w, r)
					return
				}
			}
		}
	}
}

func (a AuthorizationHandlerImpl) Authorization(next http.HandlerFunc, allowedRoles []models.Role) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contains(allowedRoles, r.Header["Role"][0]) {
			next.ServeHTTP(w, r)
		} else {
			handlers.RespondError(w, http.StatusForbidden, fmt.Errorf("you are not allowed to use this resource"))
		}
	}
}

func contains(arr []models.Role, str string) bool {
	for _, a := range arr {
		if string(a) == str {
			return true
		}
	}
	return false
}
