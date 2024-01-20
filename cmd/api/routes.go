package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// routes returns the HTTP handler for the API routes.
func (app *Config) routes() http.Handler {
	// Create a new router
	router := chi.NewRouter()

	// Configure CORS
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	router.Use(cors.Handler(corsOptions))

	// Add heartbeat middleware
	router.Use(middleware.Heartbeat("/ping"))

	router.Post("/log", app.WriteLog)

	return router
}
