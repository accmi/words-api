package main

import (
	"log"

	Config "github.com/accmi/words-api/config"
	Models "github.com/accmi/words-api/models"
	Routes "github.com/accmi/words-api/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	config := Config.DBConfig{}
	config.BuildDBConfig()
	cstring := config.DbURL()

	Config.DB, err = gorm.Open(postgres.Open(cstring), &gorm.Config{})

	if err != nil {
		log.Println("db connection error", err)
	}

	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()

	r.Run()
}
