package config

import (
	"log"
	"text/template"
)

// AppConfig holds the App Config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
