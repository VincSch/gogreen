package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMeasurementsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(FindMeasurements())
}

func CreateMeasurementEndpoint(w http.ResponseWriter, req *http.Request) {
	var measure Measurement
	_ = json.NewDecoder(req.Body).Decode(&measumrents)
	StoreMeasurement(measure)
}


func RestAPI() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/measurements", GetMeasurementsEndpoint).Methods("GET")
	router.HandleFunc("/measurement/{id}", CreateMeasurementEndpoint).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	return router
}
