package helpers

import (
	"html/template"
	"log"
	"net/http"
)

// See a web page
func RenderTemplates(w http.ResponseWriter, tmpl string) {
	ts, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
