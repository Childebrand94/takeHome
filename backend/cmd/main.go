package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable not set, defaulting to 8080")
		port = "8080"
	}
	fmt.Println("Using port:", port)

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
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
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
