package main

import (
	"encoding/json"
	"log"

	// "os"
	"divviup-client/pkg/common/db"
	"divviup-client/pkg/task_job"
	"divviup-client/pkg/common/models"

	"time"

	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ConfigProvider struct {
	DB *gorm.DB
}

func main() {
	log.Print("SCHEDULER")

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)
	d := db.Init(dbUrl)
	
	provider := &ConfigProvider{d}

	mgr, err := asynq.NewPeriodicTaskManager(
			asynq.PeriodicTaskManagerOpts{
					RedisConnOpt:               asynq.RedisClientOpt{
						Addr: viper.Get("REDIS_HOST").(string),
						Password: viper.Get("REDIS_PASSWORD").(string),
					},
					PeriodicTaskConfigProvider: provider,         // this provider object is the interface to your config source
					SyncInterval:               10 * time.Second, // this field specifies how often sync should happen
	})
	
	if err != nil {
			log.Fatal(err)
	}

	if err := mgr.Run(); err != nil {
			 log.Fatal(err)
	}
}

func (p *ConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {

	var taskJobs []models.TaskJob
	p.DB.Find(&taskJobs)
	log.Print(taskJobs)

	var configs []*asynq.PeriodicTaskConfig
	for _, cfg := range taskJobs {

			payload, err := json.Marshal(taskjob.CollectorRunPayload{TaskName:cfg.TaskName, TaskType: cfg.TaskType, DivviUpId: cfg.DivviUpId, TaskId: cfg.TaskID})
			if err != nil {
					log.Fatal(err)
			}
			
			configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.Cron, Task: asynq.NewTask("collector:run", payload)})
	}
	// return configs, nil
	return configs, nil
}
