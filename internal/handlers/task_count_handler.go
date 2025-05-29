package handlers

import (
	"task-reports/internal/database"
	"task-reports/internal/models"
	"task-reports/internal/utils"

	"github.com/gin-gonic/gin"
)

func IncrementTaskCount(c *gin.Context) {
	var input struct {
		EmployeeID uint `json:"employee_id" binding:"required"`
		TaskID     uint `json:"task_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	date := utils.TodayString()
	var taskCount models.TaskCount

	result := database.DB.Where("employee_id = ? AND task_id = ? AND date = ?", input.EmployeeID, input.TaskID, date).
		First(&taskCount)

	if result.Error != nil {
		taskCount = models.TaskCount{
			EmployeeID: input.EmployeeID,
			TaskID:     input.TaskID,
			Date:       date,
			Count:      1,
		}

		if err := database.DB.Create(&taskCount).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "message": "Ocorreu um erro ao criar o registro"})
			return
		}
	} else {
		taskCount.Count++
		if err := database.DB.Save(&taskCount).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "message": "Ocorreu um erro ao atualizar o registro"})
			return
		}
	}

	c.JSON(200, taskCount)
}
