package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-pg-mux-apis/medium/app"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("App is listening on port %s", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Fatalf("error: %v", err)
	}
}
