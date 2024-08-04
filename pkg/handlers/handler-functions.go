package handlers

import (
	"fmt"
	"log"
	"net/http"
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
	requiredLandArea, err = strconv.Atoi(requiredLandAreaRecieved)

	if err != nil && strings.Contains(err.Error(), "invalid syntax") {
		errorMap := map[string]string{}
		errorMap["integerValue"] = "You have to enter a positive value"
		render.RenderTemplate(w, r, "solar-analysis.page.tmpl", &models.TemplateData{
			Errors: errorMap,
		})
	}

	getPlanQuery := `select * from solar where land_area_minimum <= $1`
	rows, err := db.Query(getPlanQuery, requiredLandArea)
	if err != nil {
		log.Println(err)
	}

	var plan, power, modules, batteries, accessories, electricity, company, link string
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
		<a href="" class="btn btn-primary">Calculate Profit</a>
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
		err = rows.Scan(&id, &plan, &land_area_minimum, &land_area_maximum, &power, &modules, &batteries, &accessories, &electricity, &company, &link)
		if err != nil {
			log.Println(err)
		}

		land_area_minimum := strconv.Itoa(land_area_minimum)
		land_area_maximum := strconv.Itoa(land_area_maximum)

		newHtml := `
		<html>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
			<div class="card text-bg-dark" style="width: 18rem;">
				<div class="card-header" style="font-size: 1.5rem;">
					%s
				</div>
				<div class="card-body">
					<li class="list-group-item">Company: %s</li>
					<hr>
					<li class="list-group-item">Minimum Area(sqft): %s</li>
					<hr>
					<li class="list-group-item">Maxmimum Area(sqft): %s</li>
					<hr>
					<li class="list-group-item">Power: %s</li>
					<hr>
					<li class="list-group-item">Modules: %s</li>
					<hr>
					<li class="list-group-item">Batteries: %s</li>
					<hr>
					<li class="list-group-item">Accessories: %s</li>
					<hr>
					<li class="list-group-item">Annual Electricity Generated: %s</li>
					<hr>
					<li class="list-group-item"><a href="%s">Visit Site</a></li>
					<hr>
				</div>
			</div>
			<br>
			<br>
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
		</html>
		`

		fmt.Fprintf(w, newHtml, plan, company, land_area_minimum, land_area_maximum, power, modules, batteries, accessories, electricity, link)

	}
}
