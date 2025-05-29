package database

import (
	"log"
	"task-reports/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("task_control.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com banco de dados:", err)
	}

	err = DB.AutoMigrate(&models.Employee{}, &models.Task{}, &models.TaskCount{})
	if err != nil {
		log.Fatal("Erro ao migrar banco de dados:", err)
	}
}
