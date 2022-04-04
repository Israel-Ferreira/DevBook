package utils

import (
	"net/http"
	"text/template"
)

var Templates *template.Template

func LoadTemplates() {
	Templates = template.Must(template.ParseGlob("views/*.html"))
}

func RenderTemplate(rw http.ResponseWriter, templateName string, data interface{}) {
	Templates.ExecuteTemplate(rw, templateName, data)
}
