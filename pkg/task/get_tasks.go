package task

import (
	"divviup-client/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h handler) GetTasks(c *gin.Context) {
		var tasks []models.Task
		h.DB.Find(&tasks)

    c.JSON(http.StatusOK, &tasks)
}
