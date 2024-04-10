package collector

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

    routes := r.Group("/collector")
    routes.GET("/test", h.TestCollector)
    routes.POST("/manual", h.ManualRunCollector)
}