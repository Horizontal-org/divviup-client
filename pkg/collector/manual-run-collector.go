package collector

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type ManualRunCollectorRequestBody struct {
	TaskId uint `json:"task_id" binding:"required"`
	DivviUpId string `json:"divviup_id" binding:"required"`
	VdafType string `json:"type" binding:"required"`
}

func (h handler) ManualRunCollector(c *gin.Context) {
	body := ManualRunCollectorRequestBody{}
	log.Println(body)

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
	}

	ScheduledCollector(h.DB, body.VdafType, body.DivviUpId, body.TaskId)
}
