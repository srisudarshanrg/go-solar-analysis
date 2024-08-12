package handlers

import (
	"fmt"
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

	var plan, modules, batteries, accessories, electricity, company, link, cost string
	var power float64
	var id, land_area_minimum, land_area_maximum int
	html := `
	<html>
	<body style="background-color: rgb(26, 26, 26); color: #fff;">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
		<h1>Available Plans</h1>
		<br>
		<a href="/solar" class="btn btn-primary"><i class="fa-solid fa-arrow-left"></i></a>
		<br>
		<br>
		<a href="/solar-profit" class="btn btn-primary">Calculate Profit</a>
		<br>
		<br>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
		<script src="https://kit.fontawesome.com/2ce79bf423.js" crossorigin="anonymous"></script>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
	w.Header().Add("Content-Type", "text/html")

	for rows.Next() {
		err = rows.Scan(&id, &plan, &land_area_minimum, &land_area_maximum, &power, &modules, &batteries, &accessories, &electricity, &company, &link, &cost)
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

		newHtml := `
		<html>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
			<div class="row">
				<div class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
					<div class="card text-bg-dark" style="width: 18rem;">
						<div class="card-header" style="font-size: 1.5rem;">
							%s
						</div>
						<div class="card-body">
							<li class="list-group-item">Company: %s</li>
							<hr>
							<li class="list-group-item">Minimum Area(sqft): %s sqft</li>
							<hr>
							<li class="list-group-item">Maxmimum Area(sqft): %s sqft</li>
							<hr>
							<li class="list-group-item">Power: %f kVA PCU</li>
							<hr>
							<li class="list-group-item">Modules: %s</li>
							<hr>
							<li class="list-group-item">Batteries: %s</li>
							<hr>
							<li class="list-group-item">Accessories: %s</li>
							<hr>
							<li class="list-group-item">Setup Cost: %d rupees</li>
							<hr>
							<li class="list-group-item">Annual Electricity Generated: %s</li>
							<hr>
							<li class="list-group-item"><a href="%s">Visit Site</a></li>
							<hr>
						</div>
					</div>
				</div>
				<div class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
					<div class="card text-bg-dark" style="width: 18rem;">
						<div class="card-header" style="font-size: 1.5rem;">
							%s Profit Calculator
						</div>
						<div class="card-body">
							<li class="list-group-item">Your existing electricity bill (annual): %d rupees</li>
							<hr>
							<li class="list-group-item">Setup cost: %d rupees</li>
							<hr>
							<li class="list-group-item">Break Even: %d Years</li>
						</div>
					</div>
				</div>
			</div>
			
			<br>
			<br>
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
		</html>
		`

		fmt.Fprintf(w, newHtml, plan, company, land_area_minimum, land_area_maximum, power, modules, batteries, accessories, costNew, electricity, link, plan, electricityBill, costNew, time)
	}
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
