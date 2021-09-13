package handlers

import (
	"net/http"

	"github.com/markkj/simple-webapplication/pkg/config"
	"github.com/markkj/simple-webapplication/pkg/models"
	"github.com/markkj/simple-webapplication/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello, Home Page")
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello, About Page")
	stringMap := make(map[string]string)
	stringMap["test"] = "hello"
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
