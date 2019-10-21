package app

import (
	"net/http"

	u "github.com/go-pg-mux-apis/medium/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resource was not found on the server"))
		next.ServeHTTP(w, r)
	})
}
