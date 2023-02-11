package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	router := mux.NewRouter()
	router.HandleFunc("/employees", GetAllEmployees).Methods("GET")
	router.HandleFunc("/employee/{empID}", GetEmployeeByID).Methods("GET")
	router.HandleFunc("/employee", CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{empID}", UpdateEmployeeByID).Methods("PUT")
	router.HandleFunc("/employee/{empID}", DeleteEmployeeByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
