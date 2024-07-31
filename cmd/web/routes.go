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

	mux.Get("/solar-analysis", handlers.Repository.SolarAnalysis)

	mux.Get("/wind-analysis", handlers.Repository.WindAnalysis)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
