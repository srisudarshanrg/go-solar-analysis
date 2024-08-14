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

	mux.Get("/solar-profit", handlers.Repository.SolarProfit)
	mux.Post("/solar-profit", handlers.Repository.PostSolarProfit)

	mux.Get("/solar-profit-result", handlers.Repository.SolarProfitResult)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
