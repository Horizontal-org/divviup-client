package models

import (
	"time"
)

type TaskJob struct {
	ID        uint `gorm:"primarykey" json:"id"`
	TaskID uint // foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
	Cron string
	TaskName string
	TaskType string
	DivviUpId string
}