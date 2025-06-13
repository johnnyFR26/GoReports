package main

import (
	"task-reports/internal/database"
	"task-reports/internal/handlers"
	"task-reports/internal/scheduler"
	"task-reports/internal/sockets"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	go sockets.HandleMessages()

	r := gin.Default()

	scheduler.Start()

	r.GET("/ws", sockets.HandleConnections)

	r.POST("/employees", handlers.CreateEmployee)
	r.GET("/employees", handlers.GetEmployees)

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)

	r.POST("/task_counts/increment", handlers.IncrementTaskCount)

	r.Run(":8080")
}
