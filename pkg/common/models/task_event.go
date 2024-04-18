package models

import (
	"time"
)

type TaskEvent struct {
	ID        uint `gorm:"primarykey" json:"id"`
	TaskID uint // foreign key
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
	Error string `json:"error"`
	Value string `json:"value"`
}