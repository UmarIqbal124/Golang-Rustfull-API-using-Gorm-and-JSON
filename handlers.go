package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var createEmployee Employee
	json.NewDecoder(r.Body).Decode(&createEmployee)
	Database.Create(&createEmployee)
	json.NewEncoder(w).Encode(&createEmployee)

}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var getEmployees []Employee
	Database.Find(&getEmployees)
	json.NewEncoder(w).Encode(getEmployees)

}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var getEmployeeByID Employee
	Database.First(&getEmployeeByID, mux.Vars(r)["empID"])
	json.NewEncoder(w).Encode(getEmployeeByID)

}

func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var updateEmployeeByID Employee
	Database.First(&updateEmployeeByID, mux.Vars(r)["empID"])
	json.NewDecoder(r.Body).Decode(&updateEmployeeByID)
	Database.Save(&updateEmployeeByID)
	json.NewEncoder(w).Encode(updateEmployeeByID)

}

func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var deleteEmployeeByID Employee
	Database.Delete(&deleteEmployeeByID, mux.Vars(r)["empID"])
	json.NewEncoder(w).Encode("Employee deleted Successfully!")

}
