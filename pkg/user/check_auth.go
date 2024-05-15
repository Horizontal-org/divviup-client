package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Just check that token works
func (h handler) CheckAuth(c *gin.Context) {
	c.Status(http.StatusOK)
}