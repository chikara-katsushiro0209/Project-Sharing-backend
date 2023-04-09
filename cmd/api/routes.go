package main

import (
	"backend/cmd/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", controller.Home)
	mux.Get("/movies", controller.AllMovies)
	mux.Get("/users", controller.AllUsers)
	mux.Put("/users/{id}", controller.UpdateUser)

	return mux
}
