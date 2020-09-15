package main

import (
	"context"
	"fmt"
	routes "github.com/accmi/words-api/app/routes"
	config "github.com/accmi/words-api/config"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var err error

func main() {
	err = godotenv.Load(".env")

	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	configDb := config.DBConfig{}
	configDb.BuildDBConfig()
	pgString := configDb.DbURL()

	config.DB, err = pgx.Connect(context.Background(), pgString)

	if err != nil {
		log.Panicln("Problems with connection to DB", err)
	}

	r := routes.SetupRouter()


	port := os.Getenv("PORT")

	if port == "" {
		log.Panicln("port variable have not found")
		return
	}

	defer config.DB.Close(context.Background())

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
