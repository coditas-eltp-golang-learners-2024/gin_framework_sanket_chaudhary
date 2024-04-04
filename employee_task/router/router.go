package router

import (
	"github.com/gin-gonic/gin"
	"employee_task/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())   // LOG DETAILS ON CONSOLE
	router.GET("/employees", handler.GetEmployee)
	router.POST("/employees",handler.AddEmployee)
	return router
}
