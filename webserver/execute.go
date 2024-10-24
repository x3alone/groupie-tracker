package webserver

import (
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, data any) {
	// parse the web page template
	t, err := template.ParseFiles("./website/pages/home.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	// execute the templat on web page with the data serve
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
