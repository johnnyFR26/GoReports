package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `gorm:"not null"`
	TaskCounts []TaskCount
}
type Task struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	TaskCounts []TaskCount
}

type TaskCount struct {
	gorm.Model
	EmployeeID uint
	TaskID     uint
	Count      int
	Date       string `gorm:"index"`
}
