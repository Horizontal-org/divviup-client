package models

import (
	"time"
)

type TaskEvent struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	TaskID string
	Error string
	Value string
}