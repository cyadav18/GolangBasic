package handeler

import (
	"Udemy/config"
	"Udemy/pkg/model"
	"Udemy/pkg/renderer"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func SetRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

func (repository *Repository) HandleHome(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	repository.App.Session.Put(request.Context(), "remote_ip", remoteIP)
	renderer.RenderTemplate(writer, "home_tmpl.gohtml", &model.TemplateData{})
}

func (repository *Repository) HandleAbout(writer http.ResponseWriter, request *http.Request) {
	remoteIP := repository.App.Session.GetString(request.Context(),"remote_ip")
	renderer.RenderTemplate(writer, "about_tmpl.gohtml", &model.TemplateData{
		StringMap: map[string]string{"test": "hello dsadad","remote_ip":remoteIP},
	})
}
