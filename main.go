package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page Hit")
	fmt.Printf("Endppoint Hit")
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	HandleRequests()
}
