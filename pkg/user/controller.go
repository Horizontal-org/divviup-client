package user

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

    routes := r.Group("/user")
    routes.GET("/check", h.CheckAuth, middleware.TokenAuthMiddleware(db))    
    routes.POST("/login", h.Login)    
}