package taskjob

import (
	"divviup-client/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTaskJob(c *gin.Context) {

	var jobs []models.TaskJob
	h.DB.Find(&jobs)

	c.JSON(http.StatusOK, &jobs)
}