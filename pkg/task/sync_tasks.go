package task

import (
	"divviup-client/pkg/common/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

type ApiTask struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DivviUpId string
	Name string `json:"name"`
	Vdaf models.Vdaf `gorm:"embedded;embeddedPrefix:vdaf_" json:"vdaf"`
}

func (h handler) SyncTasks(c *gin.Context) {
		var apiTasks []ApiTask
		var tasks []models.Task

		GetTasksFromApi(c, &apiTasks)
	
		// Map from divviup response to orm object
		MapTasks(apiTasks, &tasks)

		// Source of truth is always DIVVIUP API		 
	 	h.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Table: "tasks", Name: "divvi_up_id"}},
			DoNothing: true,
		}).Create(&tasks)


		truthIds := GetIds(tasks)
		h.DB.Where("divvi_up_id NOT IN ?", truthIds).Delete(&models.Task{})

		var syncedTasks []models.Task
		h.DB.Find(&syncedTasks)
    c.JSON(http.StatusOK, syncedTasks)
}

func GetTasksFromApi(c *gin.Context, apiTasks *[]ApiTask) {
		req, newRerr := http.NewRequest(http.MethodGet, viper.Get("DIVVIUP_API_URL").(string) + "/accounts/" + viper.Get("DIVVIUP_ACCOUNT").(string) + "/tasks", nil)
		if newRerr != nil {
			  c.AbortWithError(http.StatusNotFound, newRerr)
        return
		}

		req.Header.Add("Accept", "application/vnd.divviup+json;version=0.1")
		req.Header.Add("Authorization", "Bearer " + viper.Get("DIVVIUP_TOKEN").(string))

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)

		//We Read the response body on the line below.
		body, err := io.ReadAll(resp.Body)
		if err != nil {
				log.Fatalln(err)
		}	 

		json.Unmarshal(body, &apiTasks)		
}

func GetIds(apiTasks []models.Task) (apiIds []string) {
	for _, apiTask := range apiTasks {
			apiIds = append(apiIds, apiTask.DivviUpId)
	}

	return
}


func MapTasks(apiTasks []ApiTask, tasks *[]models.Task) {
	for _, apiTask := range apiTasks {
		
		var newTask models.Task
		newTask.CreatedAt = apiTask.CreatedAt
		newTask.UpdatedAt = apiTask.UpdatedAt
		newTask.DivviUpId = apiTask.ID
		newTask.Vdaf = apiTask.Vdaf
		newTask.Name = apiTask.Name

		*tasks = append(*tasks, newTask)
	}
}