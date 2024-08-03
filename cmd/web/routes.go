package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/idp-project/pkg/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repository.Home)

	mux.Get("/resource-consumption", handlers.Repository.ResourceConsumption)
	mux.Post("/resource-consumption", handlers.Repository.PostResourceConsumption)

	mux.Get("/resource-production", handlers.Repository.ResourceProduction)
	mux.Post("/resource-production", handlers.Repository.PostResourceProduction)

	mux.Get("/solar", handlers.Repository.Solar)
	mux.Post("/solar", handlers.Repository.PostSolar)

<<<<<<< HEAD
	mux.Get("/wind-analysis", handlers.Repository.Wind)
=======
	mux.Get("/wind", handlers.Repository.Wind)
>>>>>>> 68bf4da4a1c010a5e394cff73844728f2a0a6322

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
