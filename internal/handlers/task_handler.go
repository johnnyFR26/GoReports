package handlers

import (
	"task-reports/internal/database"
	"task-reports/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{Name: input.Name}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(200, tasks)
}
