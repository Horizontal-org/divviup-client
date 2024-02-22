package db

import (
	"divviup-client/pkg/common/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Task{})

    return db
}