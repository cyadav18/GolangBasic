package renderer

import (
	"Udemy/config"
	"Udemy/pkg/model"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(writer http.ResponseWriter, tmpl string, data *model.TemplateData) {

	templateCash := app.TemplateCache
	fmt.Println(templateCash)
	if value, ok := templateCash[tmpl]; ok {
		buf := new(bytes.Buffer)
		_ = value.Execute(buf, data)
		_, err := buf.WriteTo(writer)
		if err != nil {
			fmt.Println("Error writing to templates")
			return
		}
	} else {
		log.Fatalln("could not create a template cache")
	}
}

func CreateTemplateCash() (map[string]*template.Template, error) {
	myCash := make(map[string]*template.Template)
	files, err := filepath.Glob("templates\\*tmpl*")
	if err != nil {
		return myCash, err
	}

	for _, file := range files {
		name := filepath.Base(file)
		parseFiles, err := template.New(name).Funcs(functions).ParseFiles(file)
		if err != nil {
			fmt.Println(err)
			return myCash, err
		}
		layouts, err := filepath.Glob("templates\\*layout*")
		if err != nil {
			return myCash, err
		}
		if len(layouts) > 0 {
			parseFiles, err = parseFiles.ParseGlob("templates\\*layout*")
			if err != nil {
				return myCash, err
			}
		}
		myCash[name] = parseFiles
	}
	fmt.Println("my cash", myCash)
	return myCash, nil
}
