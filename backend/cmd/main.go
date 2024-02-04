package main

import (
	"fmt"
	"net/http"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(middleware.Logger)
	r.Use(cors.Handler)

	// Route
	queryHandler := &handler.Query{}
	r.Post("/submit", queryHandler.Parse)

	// Start Server
	fmt.Println("Starting Server...")
	http.ListenAndServe(":3000", r)
}
