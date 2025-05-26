package main

import (
	"log"
	"net/http"
	"quotes/internal/app/quotes"
	"quotes/internal/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	service := quotes.NewService()
	handler := quotes.NewHandler(service)

	r := mux.NewRouter()
	r.Use(middleware.JSON)

	handler.RegisterRoutes(r)

	log.Println("Listening on :8080…")
	log.Fatal(http.ListenAndServe(":8080", r))
}
