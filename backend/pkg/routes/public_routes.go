package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func PublicRoutes(r *chi.Mux) {
	r.Get("/", defaultHandler)
	r.Get("/{id}", defaultHandler)

	r.Post("/api/url/create", defaultHandler)
	r.Post("/api/url/starus", defaultHandler)
}
