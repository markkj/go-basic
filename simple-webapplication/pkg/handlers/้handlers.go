package handlers

import (
	"net/http"

	"github.com/markkj/simple-webapplication/pkg/render"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello, Home Page")
	render.RenderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello, About Page")
	render.RenderTemplate(rw, "about.page.tmpl")

}
