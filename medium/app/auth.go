package app

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-pg-mux-apis/medium/models"
	u "github.com/go-pg-mux-apis/medium/utils"

	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HanderFunc(func(w http.ResponseWriter, r *http.Request) {
		// list of all endpoints that doesn't require authentication
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path // current request path

		// check if request doesn't need authentication\
		// if not, serve
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		resp := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		// if token missing
		if tokenHeader == "" {
			resp = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			resp = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			resp = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			resp = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// all successfull
		fmt.Sprintf("User %", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
