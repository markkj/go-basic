package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello, Home Page")
	renderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello, About Page")
	renderTemplate(rw, "about.page.tmpl")

}

func renderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("error parsed template: ", err)
		return
	}
}
