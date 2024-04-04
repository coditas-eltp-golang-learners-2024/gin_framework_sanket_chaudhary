package handler

import (
	"employee_task/config"
	"employee_task/model"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployee(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, model.Employee{})
	db, err := config.ConnectionWithDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Query the database to fetch employees
	rows, err := db.Query("SELECT id, name, salary, email, department FROM employees")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Clear existing employees slice
	model.Employees = []model.Employee{}

	// Populate employees slice with fetched data
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.Id, &emp.Name, &emp.Salary, &emp.Email, &emp.Department); err != nil {
			panic(err)
		}
		model.Employees = append(model.Employees, emp)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	// Return employees as JSON response
	c.JSON(http.StatusOK, model.Employees)
}

func AddEmployee(c *gin.Context) {
	var newEmployee model.Employee

	// Get database connection
	db, err := config.ConnectionWithDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Bind request body to newEmployee struct
	if err := c.BindJSON(&newEmployee); err != nil {
		panic(err)
	}

	// Execute SQL query to insert new employee
	query := "INSERT INTO employees (name, salary, email, department) VALUES (?, ?, ?, ?)"

	_, err = db.Exec(query, newEmployee.Name, newEmployee.Salary, newEmployee.Email, newEmployee.Department)
if err != nil {
   panic(err)
}


	// Append newEmployee to model.Employees slice
	model.Employees = append(model.Employees, newEmployee)

	// Return newly added employee as JSON response
	c.JSON(http.StatusCreated, newEmployee)
}

