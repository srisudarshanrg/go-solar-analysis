package handlers

import (
	"database/sql"
	"net/http"

	"github.com/srisudarshanrg/go-solar-analysis/pkg/config"
	"github.com/srisudarshanrg/go-solar-analysis/pkg/models"
	"github.com/srisudarshanrg/go-solar-analysis/pkg/render"
)

var Repository HandlerAccess
var db *sql.DB

// HandlerAccess holds the application configuration
type HandlerAccess struct {
	app *config.AppConfig
}

// SetUpAppConfig sets the application configuration for the handlers package
func SetUpAppConfig(a *config.AppConfig) *HandlerAccess {
	return &HandlerAccess{
		app: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *HandlerAccess) {
	Repository = *r
}

// DBAccess provides access for the database to the handlers package
func DBAccess(database *sql.DB) {
	db = database
}

// Home is the handler for the home page
func (a *HandlerAccess) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// ResourceConsumption is the handler for the resource consumption page
func (a *HandlerAccess) ResourceConsumption(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resource-consumption.page.tmpl", &models.TemplateData{})
}

// PostResourceConsumption handles the posted form from the resource consumption page
func (a *HandlerAccess) PostResourceConsumption(w http.ResponseWriter, r *http.Request) {
	PostResourceConsumptionFunction(w, r)
}

// ResourceProduction is the handler for the resource production page
func (a *HandlerAccess) ResourceProduction(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resource-production.page.tmpl", &models.TemplateData{})
}

// PostResourceProduction handles the posted form from the resource production page
func (a *HandlerAccess) PostResourceProduction(w http.ResponseWriter, r *http.Request) {
	PostResourceProductionFunction(w, r)
}

// Solar is the handler for the solar page
func (a *HandlerAccess) Solar(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{})
}

// PostSolar handles the posted form from the solar page
func (a *HandlerAccess) PostSolar(w http.ResponseWriter, r *http.Request) {
	PostSolarFunction(w, r)
}

func (a *HandlerAccess) SolarResult(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-result.page.tmpl", &models.TemplateData{})
}

// SolarProfit is the handler for the solar profit page
func (a *HandlerAccess) SolarProfit(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-profit.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) PostSolarProfit(w http.ResponseWriter, r *http.Request) {
	PostSolarProfitFunction(w, r)
}

func (a *HandlerAccess) SolarProfitResult(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-profit-result.page.tmpl", &models.TemplateData{})
}
