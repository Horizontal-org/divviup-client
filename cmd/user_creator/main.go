package main

import (
	"divviup-client/pkg/common/db"
	"divviup-client/pkg/common/models"
	"log"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// first user then password
func main() {
	log.Println("USER_CREATOR")
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	
	username := os.Args[1]
	password := os.Args[2]

	dbUrl := viper.Get("DB_URL").(string)
	d := db.Init(dbUrl)
	
	var userFound models.User
	d.Where("username=?", username).Find(&userFound)

	if userFound.ID != 0 {
		log.Println("ERROR 422")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERROR 422")
		return
	}

	user := models.User{
		Username: username,
		Password: string(passwordHash),
	}

	d.Create(&user)
	log.Println("SUCCESS")
}

