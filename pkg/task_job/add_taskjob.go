package taskjob

import (
	"divviup-client/pkg/common/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AddTaskJobRequestBody struct {
	TaskId uint `json:"task_id" binding:"required"`
}

func (h handler) AddTaskJob(c *gin.Context) {
	body := AddTaskJobRequestBody{}
	log.Println(body)

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
	}
	
	var task models.Task
	h.DB.First(&task, body.TaskId)
	task.Starred = true	
	h.DB.Save(&task)

	taskJob := models.TaskJob{
		//TODO CHaNGE WITH REAL DATA
		Cron: "0 */12 * * *",
		TaskName: task.Name,
		TaskType: task.Vdaf.Type,
		TaskID: task.ID,
		DivviUpId: task.DivviUpId,
	}

	h.DB.Create(&taskJob)

	var tasks []models.Task
	h.DB.Find(&tasks)

	c.JSON(http.StatusOK, &tasks)
}