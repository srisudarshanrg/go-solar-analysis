package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/srisudarshanrg/idp-project/pkg/config"
	"github.com/srisudarshanrg/idp-project/pkg/database"
	"github.com/srisudarshanrg/idp-project/pkg/handlers"
	"github.com/srisudarshanrg/idp-project/pkg/render"

	"github.com/go-chi/chi"
)

var app config.AppConfig
var db *sql.DB

const portNumber = ":4000"

func main() {
	var err error
	app.UseCache = false

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = templateCache

	repository := handlers.SetUpAppConfig(&app)
	handlers.NewHandlers(repository)

	render.SetAppConfig(&app)

	// create database connection
	db, err = database.CreateDatabaseConnection()
	if err != nil {
		log.Println(err)
	}

	handlers.DBAccess(db)

	defer db.Close()

	// run the routes
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repository.Home)

	mux.Get("/resource-consumption", handlers.Repository.ResourceConsumption)
	mux.Post("/resource-consumption", handlers.Repository.PostResourceConsumption)

	mux.Get("/resource-production", handlers.Repository.ResourceProduction)
	mux.Post("/resource-production", handlers.Repository.PostResourceProduction)

	mux.Get("/solar", handlers.Repository.Solar)
	mux.Post("/solar", handlers.Repository.PostSolar)

	mux.Get("/solar-result", handlers.Repository.SolarResult)

	mux.Get("/solar-profit", handlers.Repository.SolarProfit)
	mux.Post("/solar-profit", handlers.Repository.PostSolarProfit)

	mux.Get("/solar-profit-result", handlers.Repository.SolarProfitResult)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
