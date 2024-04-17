package collector

import (
	"divviup-client/pkg/common/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type ManualRunCollectorRequestBody struct {
	TaskId uint `json:"task_id" binding:"required"`
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

	arguments := CollectorArguments{
		// make manifest env variable
		Manifest: "/home/juan/code/janus-0.7.0-prerelease-2/Cargo.toml",
		TaskId: strconv.FormatUint(uint64(body.TaskId), 10),
		VdafType: body.VdafType,
		LeaderUrl: "https://dap-07-1.api.divviup.org/",
	}


	collectorOut, collectorSuccess := RunCollector(&arguments)


	var taskEvent models.TaskEvent
	taskEvent.TaskID = body.TaskId
	
	if (collectorSuccess) {
		taskEvent.Value = collectorOut
	} else {
		taskEvent.Error = collectorOut
	}

	h.DB.Create(&taskEvent)


	log.Println("Output: ", collectorOut)
	log.Println("Success: ", collectorSuccess)
	
}
