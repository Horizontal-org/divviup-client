package taskjob

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/taskjob")
    routes.POST("/add", h.AddTaskJob)
    routes.POST("/delete", h.DeleteTaskJob)
}