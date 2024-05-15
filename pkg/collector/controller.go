package collector

import (
	"divviup-client/pkg/common/middleware"

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

    routes := r.Group("/collector", middleware.TokenAuthMiddleware(db))
    routes.GET("/test", h.TestCollector)
    routes.POST("/manual", h.ManualRunCollector)
}