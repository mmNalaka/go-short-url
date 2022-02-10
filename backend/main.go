package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mmnalaka/go-short-url/pkg/middlewares"
	"github.com/mmnalaka/go-short-url/pkg/routes"
	"github.com/mmnalaka/go-short-url/platfrom/cache"
	"github.com/mmnalaka/go-short-url/platfrom/database"
)

func main() {
	r := chi.NewRouter()

	// connect to database
	database.Init()
	database.RunMigrations()

	// connect cache
	cache := cache.Pool.Get()
	defer cache.Close()

	// Middlewares
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middlewares.CORSConfig())

	r.Get("/set/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		value, err := cache.Do("SET", id, id)
		if err != nil {
			panic(err)
		}

		rw.Write([]byte(value.(string)))
	})

	r.Get("/get/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		value, _ := cache.Do("GET", id)
		fmt.Printf("%s \n", value)
	})

	// Routes
	routes.PublicRoutes(r)

	http.ListenAndServe(":4000", r)
}
