package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Get("/example-endpoint", app.ExampleEndpoint)
	return router
}
