package task

import (
	"divviup-client/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h handler) GetStarredTasks(c *gin.Context) {
		var tasks []models.Task
		h.DB.Where("starred = ?", true).Find(&tasks)

    c.JSON(http.StatusOK, &tasks)
}
