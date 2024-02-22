package models

import (
	"time"
)

type Vdaf struct {
	Type string `json:"type"`
}

type Task struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string `json:"name"`
	Vdaf Vdaf `gorm:"embedded;embeddedPrefix:vdaf_"`
	CollectorCredentialId string `json:"collector_credential_id"`
}