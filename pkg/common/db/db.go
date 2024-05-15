package db

import (
	"divviup-client/pkg/common/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(mysql.Open(url), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Task{}, &models.TaskJob{}, &models.TaskEvent{}, &models.User{})

    return db
}