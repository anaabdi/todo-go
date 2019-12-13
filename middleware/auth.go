package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/anaabdi/todo-go/repository"

	"github.com/dgrijalva/jwt-go"

	"github.com/anaabdi/todo-go/handler"
	"github.com/anaabdi/todo-go/helper"
)

// HTTP middleware setting a value on the request context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("err: missing auth token")
			handler.Respond(w, http.StatusUnauthorized, nil)
			return
		}

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		if bearerToken == "" {
			log.Printf("err: missing auth token")
			handler.Respond(w, http.StatusUnauthorized, nil)
			return
		}

		log.Println(bearerToken)

		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				log.Printf("err: different method, possibly tempered token")
				handler.Respond(w, http.StatusUnauthorized, nil)
				return nil, errors.New("different method, possibly tempered token")
			}

			return helper.JWTVerifyingKey, nil
		})

		if err != nil {
			log.Printf("err: failed to parse token")
			handler.Respond(w, http.StatusUnauthorized, nil)
			return
		}

		if !token.Valid {
			log.Printf("err: invalid token")
			handler.Respond(w, http.StatusUnauthorized, nil)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("err: invalid token")
			handler.Respond(w, http.StatusUnauthorized, nil)
			return
		}

		username := claims["uname"].(string)

		user, err := repository.GetUserByUsername(username)
		if err != nil {
			log.Printf("err: failed to get user")
			handler.Respond(w, http.StatusBadRequest, nil)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
