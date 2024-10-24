package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	help "tools/tools"
)

func main() {
	Port := "localhost:8080"
	var err error
	if len(os.Args) > 1 {
		fmt.Println("Error : No args needed to run the program")
		return
	}

	help.Data.Cards, err = help.FetchArtistData("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		log.Printf("Error fetching artist data: %v", err)
		return
	}

	http.HandleFunc("/static/", help.ServeHandle)
	http.HandleFunc("/", help.Index)
	http.HandleFunc("/bandsinfo", help.Bandinfo)
	http.HandleFunc("/filter", help.FilterHandler)

	fmt.Println("Server is running at http://" + Port)
	err = http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
