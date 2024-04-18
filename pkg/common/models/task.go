package models

import (
	"time"
)

type Vdaf struct {
	Type string `json:"type"`
}

type Task struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DivviUpId string `gorm:"unique"`
	Name string `json:"name"`
	Vdaf Vdaf `gorm:"embedded;embeddedPrefix:vdaf_" json:"vdaf"`
	Starred bool `json:"starred"`
	// TODO add if batch or time based 
	TaskEvents []TaskEvent
	TaskJob TaskJob
}