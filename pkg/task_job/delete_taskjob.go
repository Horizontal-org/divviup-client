package taskjob

import (
	"divviup-client/pkg/common/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type DeleteTaskJobRequestBody struct {
	TaskId uint `json:"task_id" binding:"required"`
}

func (h handler) DeleteTaskJob(c *gin.Context) {
	body := DeleteTaskJobRequestBody{}
	log.Println(body)

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
	}
	
	var task models.Task
	h.DB.Preload("TaskJob").First(&task, body.TaskId)
	task.Starred = false	
	h.DB.Save(&task)

	if task.TaskJob.ID != 0 {
		h.DB.Delete(&models.TaskJob{}, task.TaskJob.ID)
	}


	var tasks []models.Task
	h.DB.Find(&tasks)

	c.JSON(http.StatusOK, &tasks)
}