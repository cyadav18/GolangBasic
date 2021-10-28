package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
