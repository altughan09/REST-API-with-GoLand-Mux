package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Country Struct (Model)
type Country struct {
	ID           string `json:"id"`
	NAME         string `json:"name"`
	CURRENCY     string `json:"currency"`
	CURRENCYCODE string `json:"currencyCode"`
}

var countries []Country

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func main() {
	//Init router
	r := mux.NewRouter()

	//Mock data
	countries = append(countries, Country{ID: "254", NAME: "UNITED STATES OF AMERICA", CURRENCY: "US DOLLAR", CURRENCYCODE: "USD"})
	//Route handlers - Endpoints
	r.HandleFunc("/api/countries", getCountries).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))
}
