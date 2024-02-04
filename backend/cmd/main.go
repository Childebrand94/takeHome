package main

import (
	"fmt"
	"net/http"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	queryHandler := &handler.Query{}
	r.Post("/submit", queryHandler.Parse)
	fmt.Println("Starting Server...")
	http.ListenAndServe(":3000", r)
}
