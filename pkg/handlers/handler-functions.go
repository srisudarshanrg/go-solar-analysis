package handlers

import (
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/srisudarshanrg/go-solar-analysis/pkg/models"
	"github.com/srisudarshanrg/go-solar-analysis/pkg/render"
)

// PostResourceConsumptionFunction is the functionality for the resource consumption feature
func PostResourceConsumptionFunction(w http.ResponseWriter, r *http.Request) {
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

func PostResourceProductionFunction(w http.ResponseWriter, r *http.Request) {
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

// PostSolarFunction is the functionality for the solar analysis feature
func PostSolarFunction(w http.ResponseWriter, r *http.Request) {
	var err error
	var requiredLandArea int

	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	requiredLandAreaRecieved := r.Form.Get("land_area")
	requiredPowerRecieved := r.Form.Get("power")
	electricityBillRecieved := r.Form.Get("billCurrent")

	electricityBill, err := strconv.Atoi(electricityBillRecieved)
	if err != nil {
		log.Println(err)
	}

	var requiredPowerRecievedFloat float64

	requiredPowerRecievedFloat, _ = strconv.ParseFloat(requiredPowerRecieved, 64)
	requiredLandArea, err = strconv.Atoi(requiredLandAreaRecieved)

	if err != nil && strings.Contains(err.Error(), "invalid syntax") {
		errorMap := map[string]string{}
		errorMap["invalidFormat"] = "Enter details in format specified"
		render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{
			Errors: errorMap,
		})
	}

	getPlanQuery := `select * from solar where land_area_minimum <= $1 and power >= $2`
	result, _ := db.Exec(getPlanQuery, requiredLandArea, requiredPowerRecievedFloat)
	affected, _ := result.RowsAffected()

	if affected == 0 {
		errorMap := map[string]string{}
		errorMap["invalidFormat"] = "No such plan is available. Perhaps you would like to check the available plans first?"
		render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{
			Errors: errorMap,
		})
	}

	rows, err := db.Query(getPlanQuery, requiredLandArea, requiredPowerRecievedFloat)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var plan, modules, batteries, accessories, electricity, company, link, cost, typePlan string
	var power float64
	var id, land_area_minimum, land_area_maximum int

	type Plan struct {
		PlanName        string
		Company         string
		PlanType        string
		LandAreaMin     string
		LandAreaMax     string
		Power           float64
		Modules         string
		Batteries       string
		Accessories     string
		Cost            int
		Electricity     string
		Link            string
		Time            int
		ElectricityBill int
	}

	var completePlanList []Plan

	for rows.Next() {
		err = rows.Scan(&id, &plan, &land_area_minimum, &land_area_maximum, &power, &modules, &batteries, &accessories, &electricity, &company, &link, &cost, &typePlan)
		if err != nil {
			log.Println(err)
		}

		land_area_minimum := strconv.Itoa(land_area_minimum)
		land_area_maximum := strconv.Itoa(land_area_maximum)

		costNew, err := strconv.Atoi(cost)
		if err != nil {
			log.Println(err)
		}

		var time int

		for i := 1; electricityBill <= costNew; i++ {
			if electricityBill*i >= costNew {
				time = i
				break
			} else {
				continue
			}
		}

		structure := Plan{
			PlanName:        plan,
			Company:         company,
			PlanType:        typePlan,
			LandAreaMin:     land_area_minimum,
			LandAreaMax:     land_area_maximum,
			Power:           power,
			Modules:         modules,
			Batteries:       batteries,
			Accessories:     accessories,
			Cost:            costNew,
			Electricity:     electricity,
			Link:            link,
			Time:            time,
			ElectricityBill: electricityBill,
		}

		completePlanList = append(completePlanList, structure)
	}

	data := map[string]interface{}{}
	data["solarPlans"] = completePlanList

	render.RenderTemplate(w, r, "solar-result.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

func PostSolarProfitFunction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	plan := r.Form.Get("plan")
	currentCostRecieved := r.Form.Get("existingCost")
	currentCost, err := strconv.Atoi(currentCostRecieved)
	if err != nil {
		log.Println(err)
	}

	getCostQuery := `select cost from solar where plan=$1`
	rows, err := db.Query(getCostQuery, plan)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var solarCostRecieved string

	for rows.Next() {
		err = rows.Scan(&solarCostRecieved)
		if err != nil {
			log.Println(err)
		}
	}

	solarCost, err := strconv.Atoi(solarCostRecieved)
	if err != nil {
		log.Println(err)
	}

	var time int

	for i := 1; currentCost <= solarCost; i++ {
		if currentCost*i >= solarCost {
			time = i
			break
		} else {
			continue
		}
	}

	type Details struct {
		CurrentCost int
		SolarCost   int
		ProfitTime  int
	}

	profitDetails := Details{
		CurrentCost: currentCost,
		SolarCost:   solarCost,
		ProfitTime:  time,
	}

	data := map[string]interface{}{}
	data["profitDetails"] = profitDetails

	render.RenderTemplate(w, r, "solar-profit-result.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// SendEmail sends an email to the given email
func SendEmail(from string, to []string, message []byte, password string) error {
	// smtp server configurations
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
