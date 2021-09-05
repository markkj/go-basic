package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/markkj/simple-webapplication/pkg/config"
	"github.com/markkj/simple-webapplication/pkg/models"
)

var fuctions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = RenderTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(rw)

	if err != nil {
		log.Fatal(err)
		return
	}
}
func RenderTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(fuctions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err

		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
