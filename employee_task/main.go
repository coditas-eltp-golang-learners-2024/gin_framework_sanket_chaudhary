package main

import (
	// "employee_task/config"
	"employee_task/router"
)

func main() {
	r := router.SetupRouter()
	// config.ConnectionWithDb()

	r.Run("localhost:8080")
}
