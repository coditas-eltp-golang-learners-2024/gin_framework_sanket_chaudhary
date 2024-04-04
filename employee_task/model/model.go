package model


type Employee struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Salary float64  `json:"salary"`
	Email     string  `json:"email"`
	Department string  `json:"department"`
}

var Employees []Employee = []Employee{}
  