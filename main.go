package main

import (
	"task-reports/internal/database"
	"task-reports/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.POST("/employees", handlers.CreateEmployee)
	r.GET("/employees", handlers.GetEmployees)

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)

	r.Run(":8080")
}
