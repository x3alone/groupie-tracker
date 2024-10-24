package tools

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

func ServeHandle(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("." + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") || err != nil {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}
	file := http.Dir("static")
	fs := http.StripPrefix("/static/", http.FileServer(file))
	fs.ServeHTTP(w, r)
}

func ExecuteTemplate(w http.ResponseWriter, data any) {
	// parse the web page template
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ExecuteError(w, "Status Internal Server Error", "500")
		return
	}
	// execute the templat on web page with the data serve
	err = t.Execute(w, data)
	if err != nil {
		ExecuteError(w, "Status Internal Server Error", "500")
		return
	}
}

func ExecuteError(w http.ResponseWriter, msg string, stat string) {
	data := struct {
		NumError     string
		MessageError string
	}{
		NumError:     stat,
		MessageError: msg,
	}
	t, err := template.ParseFiles("templates/PageError.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
