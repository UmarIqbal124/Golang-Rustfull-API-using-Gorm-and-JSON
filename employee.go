package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string  `json:"empName"`
	EmpSalary float64 `json:"empSalary"`
	EmpEmail  string  `json:"empEmail"`
}
