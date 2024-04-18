package models

import (
	"time"
)

type TaskJob struct {
	ID        uint `gorm:"primarykey" json:"id"`
	TaskID uint // foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
	Cron string `json:"cron"`
	TaskName string `json:"task_name"`
 	TaskType string `json:"task_type"`
	DivviUpId string
}