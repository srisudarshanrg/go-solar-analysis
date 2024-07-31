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

	var country, oil, electricity, coal, natural_gas, biofuel string
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
		Country     string
		Oil         string
		Electricity string
		Coal        string
		NaturalGas  string
		Biofuel     string
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

func (a *HandlerAccess) ResourceProduction(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "resource-production.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) PostResourceProduction(w http.ResponseWriter, r *http.Request) {

}

func (a *HandlerAccess) SolarAnalysis(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) WindAnalysis(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "wind-analysis.page.tmpl", &models.TemplateData{})
}
