package handlers

import (
	"task-reports/internal/database"
	"task-reports/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	employee := models.Employee{Name: input.Name}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(500, gin.H{"error": "Ocorreu um erro ao criar o funcionário", "message": err.Error()})
		return
	}

	c.JSON(201, employee)
}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	database.DB.Find(&employees)
	c.JSON(200, employees)
}
