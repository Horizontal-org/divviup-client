package main

import (
	"context"
	"divviup-client/pkg/collector"
	"divviup-client/pkg/common/db"
	"divviup-client/pkg/task_job"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	log.Println("WORKER")

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)
	DB = db.Init(dbUrl)
	
	srv := asynq.NewServer(
			asynq.RedisClientOpt{
				Addr: viper.Get("REDIS_HOST").(string),
				Password: viper.Get("REDIS_PASSWORD").(string),
			},
			asynq.Config{Concurrency: 10},
	)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
			log.Fatal(err)
	}
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "test:test":
			var p taskjob.TestTaskPayload
			if err := json.Unmarshal(t.Payload(), &p); err != nil {
					return err
			}
			log.Printf(" [*] TEST TASK %d", p.TaskType)
	case "collector:run":
			var p taskjob.CollectorRunPayload
			if err := json.Unmarshal(t.Payload(), &p); err != nil {
				return err
			}
			log.Printf(" [*] RUN COLLECTOR %d", p.TaskName, p.TaskType, p.DivviUpId, p.TaskId)
			collector.ScheduledCollector(DB, p.TaskType, p.DivviUpId, p.TaskId)
	default:
			return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}
