package main

import (
	"log"

	config "github.com/accmi/words-api/config"
	models "github.com/accmi/words-api/models"
	routes "github.com/accmi/words-api/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	configDb := config.DBConfig{}
	configDb.BuildDBConfig()
	cstring := configDb.DbURL()

	config.DB, err = gorm.Open(postgres.Open(cstring), &gorm.Config{})

	if err != nil {
		log.Println("db connection error", err)
	}

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()

	r.Run()
}
