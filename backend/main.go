package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mmnalka/go-short-url/pkg/routes"
)

func main() {
	r := chi.NewRouter()

	routes.PublicRoutes(r)

	http.ListenAndServe(":3000", r)
}
