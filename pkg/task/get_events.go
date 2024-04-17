package task

import (
	"divviup-client/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetTaskEventsQuery struct {
	id uint
} 

func (h handler) GetTaskEvents(c *gin.Context) {
		var taskEvents []models.TaskEvent

		taskId := c.Query("id")

		h.DB.Where("task_id = ?", taskId).Find(&taskEvents)
    c.JSON(http.StatusOK, &taskEvents)
}
