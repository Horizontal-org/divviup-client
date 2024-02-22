package tasks

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

func (h handler) SyncTasks(c *gin.Context) {
		var apiTasks []models.Task
		GetTasksFromApi(c, &apiTasks)

		// Source of truth is always DIVVIUP API		 
	 	h.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&apiTasks)

		apiIds := GetIds(apiTasks)
		h.DB.Where("id NOT IN ?", apiIds).Delete(&models.Task{})

    c.JSON(http.StatusOK, &apiTasks)
}

func GetTasksFromApi(c *gin.Context, apiTasks *[]models.Task) {
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

		// Parse to models
		json.Unmarshal(body, &apiTasks)
}

func GetIds(apiTasks []models.Task) (apiIds []string) {
	for _, apiTask := range apiTasks {
			apiIds = append(apiIds, apiTask.ID)
	}

	return
}


