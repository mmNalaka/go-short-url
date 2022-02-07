package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mmnalka/go-short-url/pkg/middlewares"
	"github.com/mmnalka/go-short-url/pkg/routes"
)

func main() {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middlewares.CORSConfig())

	// Routes
	routes.PublicRoutes(r)

	http.ListenAndServe(":4000", r)
}
