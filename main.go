package main

import (
	"context"
	config "github.com/accmi/words-api/config"
	routes "github.com/accmi/words-api/routes"
	"github.com/jackc/pgx/v4"
	"log"
)

var err error

func main() {
	configDb := config.DBConfig{}
	configDb.BuildDBConfig()
	pgString := configDb.DbURL()

	config.DB, err = pgx.Connect(context.Background(), pgString)

	if err != nil {
		log.Panicln("Problems with connection to DB", err)
	}
	defer config.DB.Close(context.Background())

	r := routes.SetupRouter()

	r.Run()
}
