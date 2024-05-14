package models

import (
	"time"
)

type TaskEvent struct {
	ID        uint `gorm:"primarykey" json:"id"`
	TaskID uint // foreign key
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	
	Value string `json:"value"`
	Success bool `json:"success"`
	Output string `json:"output"`
}