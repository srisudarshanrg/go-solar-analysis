package config

import (
	"database/sql"
	"text/template"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	Database      *sql.DB
	UseCache      bool
}
