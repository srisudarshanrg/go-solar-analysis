package handlers

import (
	"database/sql"
	"net/http"

	"github.com/srisudarshanrg/idp-project/pkg/config"
	"github.com/srisudarshanrg/idp-project/pkg/models"
	"github.com/srisudarshanrg/idp-project/pkg/render"
)

var Repository HandlerAccess
var db *sql.DB

type HandlerAccess struct {
	app *config.AppConfig
}

func SetUpAppConfig(a *config.AppConfig) *HandlerAccess {
	return &HandlerAccess{
		app: a,
	}
}

func NewHandlers(r *HandlerAccess) {
	Repository = *r
}

func DBAccess(database *sql.DB) {
	db = database
}

func (a *HandlerAccess) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) ResourceConsumption(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resource-consumption.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) ResourceProduction(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resource-production.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) SolarAnalysis(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) WindAnalysis(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "wind-analysis.page.tmpl", &models.TemplateData{})
}
