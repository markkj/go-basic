package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markkj/simple-webapplication/pkg/config"
	"github.com/markkj/simple-webapplication/pkg/handlers"
	"github.com/markkj/simple-webapplication/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.RenderTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application running on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
