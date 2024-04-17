package main

import (
	"divviup-client/pkg/collector"
	"divviup-client/pkg/common/db"
	"divviup-client/pkg/task"
	taskjob "divviup-client/pkg/task_job"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)	
    dbUrl := viper.Get("DB_URL").(string)
    d := db.Init(dbUrl)
		
    router := gin.Default()
    router.Use(cors.Default())
    
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "OK",
        })
    })

    router.GET("/os", func(c *gin.Context) {
        c.String(200, runtime.GOOS)
    })

    // REGISTER ROUTES 
	task.RegisterRoutes(router, d)
    taskjob.RegisterRoutes(router, d)
    collector.RegisterRoutes(router, d)

    router.Run(port)
}
