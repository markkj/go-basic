package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	InfoLog       *log.Logger
	TemplateCache map[string]*template.Template
}
