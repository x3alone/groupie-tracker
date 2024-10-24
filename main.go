package main

import (
	"log"
	"net/http"

	"GTapi/tracker"
	"GTapi/webserver"
)

var API = "https://groupietrackers.herokuapp.com/api"

func main() {
	port := ":8080"

	tracker.APiProcess(API)

	// serving style
	http.HandleFunc("/style/", webserver.ServeHandle)

	// handle web functions
	http.HandleFunc("/", webserver.HomeHandle)

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")
	// lanche the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
