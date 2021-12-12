package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/komalshah1987/GolangAPI/route"
	"github.com/komalshah1987/GolangAPI/services"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page Hit")
	fmt.Printf("Endppoint Hit")
}

func HandleRequests() {
	go services.InitializeData()
	r := route.NewRoute()
	http.HandleFunc("/", HomePage)
	go services.InitializeData()
	r := route.NewRoute()
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	HandleRequests()
}
