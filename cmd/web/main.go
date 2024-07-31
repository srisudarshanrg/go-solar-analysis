package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/srisudarshanrg/idp-project/pkg/config"
	"github.com/srisudarshanrg/idp-project/pkg/database"
	"github.com/srisudarshanrg/idp-project/pkg/handlers"
	"github.com/srisudarshanrg/idp-project/pkg/render"
)

var app config.AppConfig
var db *sql.DB

const portNumber = ":4040"

func main() {
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
	db, err := database.CreateDatabaseConnection()
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
