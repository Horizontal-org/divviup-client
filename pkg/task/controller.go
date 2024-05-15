package task

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

    routes := r.Group("/task", middleware.TokenAuthMiddleware(db))
    routes.GET("/sync", h.SyncTasks)
    routes.GET("/starred", h.GetStarredTasks)
    routes.GET("/events", h.GetTaskEvents)
    routes.GET("/", h.GetTasks)
}