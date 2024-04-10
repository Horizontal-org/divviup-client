package models

import (
	"time"
)

type CollectorWorker struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	LastRun time.Time
	RunInterval string
	TaskID string
}