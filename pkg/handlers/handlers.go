package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/srisudarshanrg/idp-project/pkg/config"
	"github.com/srisudarshanrg/idp-project/pkg/models"
	"github.com/srisudarshanrg/idp-project/pkg/render"
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
	var err error
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	countryEntered := r.Form.Get("country")
	countryEntered = strings.ToLower(countryEntered)

	// check if country exist
	getCountryQuery := `select * from resource_consumption where lower(country)=$1`
	result, err := db.Exec(getCountryQuery, countryEntered)
	if err != nil {
		log.Println(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		errorMap := map[string]string{}

		errorMap["consumptionCountryNotFound"] = "Data not available for this country."

		render.RenderTemplate(w, r, "resource-consumption.page.tmpl", &models.TemplateData{
			Errors: errorMap,
		})

		return
	}

	var country, oil, electricity, coal, natural_gas, biofuel interface{}
	var id int
	var created, updated interface{}

	row, err := db.Query(getCountryQuery, countryEntered)
	if err != nil {
		log.Println(err)
	}

	defer row.Close()

	for row.Next() {
		err = row.Scan(&id, &country, &oil, &electricity, &coal, &natural_gas, &biofuel, &created, &updated)

		if err != nil {
			log.Println(err)
		}
	}

	type ConsumptionDetails struct {
		Country     interface{}
		Oil         interface{}
		Electricity interface{}
		Coal        interface{}
		NaturalGas  interface{}
		Biofuel     interface{}
	}

	countryConsumption := ConsumptionDetails{
		Country:     country,
		Oil:         oil,
		Electricity: electricity,
		Coal:        coal,
		NaturalGas:  natural_gas,
		Biofuel:     biofuel,
	}

	dataMap := map[string]interface{}{}
	dataMap["countryConsumption"] = countryConsumption

	render.RenderTemplate(w, r, "resource-consumption.page.tmpl", &models.TemplateData{
		Data: dataMap,
	})
}

// ResourceProduction is the handler for the resource production page
func (a *HandlerAccess) ResourceProduction(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resource-production.page.tmpl", &models.TemplateData{})
}

// PostResourceProduction handles the posted form from the resource production page
func (a *HandlerAccess) PostResourceProduction(w http.ResponseWriter, r *http.Request) {
	var err error
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	countryEntered := r.Form.Get("country")
	countryEntered = strings.ToLower(countryEntered)

	yearEntered := r.Form.Get("year")

	getCountryQuery := `select * from resource_production where lower(country)=$1 and year=$2`
	result, err := db.Exec(getCountryQuery, countryEntered, yearEntered)

	if err != nil {
		log.Println(err)
	}

	rowNumber, _ := result.RowsAffected()
	if rowNumber == 0 {
		errorMap := map[string]string{}
		errorMap["productionCountryNotFound"] = "Data not available for this country or year"

		render.RenderTemplate(w, r, "resource-production.page.tmpl", &models.TemplateData{
			Errors: errorMap,
		})

		return
	}

	rows, err := db.Query(getCountryQuery, countryEntered, yearEntered)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var country, code, year, gas_production, coal_production, oil_production interface{}
	var id int
	var created, updated interface{}

	for rows.Next() {
		err = rows.Scan(&id, &country, &code, &year, &gas_production, &coal_production, &oil_production, &created, &updated)

		if err != nil {
			log.Println(err)
		}
	}

	type ProductionDetails struct {
		Country        interface{}
		Code           interface{}
		Year           interface{}
		GasProduction  interface{}
		CoalProduction interface{}
		OilProduction  interface{}
	}

	countryProduction := ProductionDetails{
		Country:        country,
		Code:           code,
		Year:           year,
		GasProduction:  gas_production,
		CoalProduction: coal_production,
		OilProduction:  oil_production,
	}

	data := map[string]interface{}{}
	data["countryProduction"] = countryProduction

	render.RenderTemplate(w, r, "resource-production.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Solar is the handler for the solar page
func (a *HandlerAccess) Solar(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{})
}

// PostSolar handles the posted form from the solar page
func (a *HandlerAccess) PostSolar(w http.ResponseWriter, r *http.Request) {
	PostSolarFunction(w, r)
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
