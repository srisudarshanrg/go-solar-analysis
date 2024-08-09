package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/srisudarshanrg/idp-project/pkg/config"
	"github.com/srisudarshanrg/idp-project/pkg/models"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	templateCache := map[string]*template.Template{}
	var err error

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	template, check := templateCache[tmpl]
	if !check {
		log.Println("Template not found in template cache")
	}

	err = template.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all template files
	files, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// parse all template files and add to template cache
	for _, file := range files {
		templateSet, err := template.New(filepath.Base(file)).ParseFiles(file)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("Error while getting layout files")
			return templateCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println(err)
				return templateCache, err
			}
		}

		templateCache[filepath.Base(file)] = templateSet
	}

	// return template cache with no errors
	return templateCache, nil
}
