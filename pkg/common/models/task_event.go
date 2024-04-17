package models

import (
	"time"
)

type TaskEvent struct {
	ID        uint `gorm:"primarykey" json:"id"`
	TaskID uint // foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
	Error string
	Value string
}