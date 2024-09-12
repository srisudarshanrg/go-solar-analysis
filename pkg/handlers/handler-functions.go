package handlers

import (
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/srisudarshanrg/idp-project/pkg/models"
	"github.com/srisudarshanrg/idp-project/pkg/render"
)

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

	log.Println(data["solarPlans"])

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

	annualElectricity := currentCost * 12
	for i := 1; annualElectricity <= solarCost; i++ {
		if annualElectricity*i >= solarCost {
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
		CurrentCost: annualElectricity,
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
