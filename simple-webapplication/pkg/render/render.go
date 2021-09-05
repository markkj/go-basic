package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("error parsed template: ", err)
		return
	}
}
