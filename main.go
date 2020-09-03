package main

import (
	"context"
	config "github.com/accmi/words-api/config"
	routes "github.com/accmi/words-api/routes"
	"github.com/jackc/pgx/v4"
	//"github.com/golang-migrate/migrate/v4"
	//"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
)

var err error

func main() {
	configDb := config.DBConfig{}
	configDb.BuildDBConfig()
	pgString := configDb.DbURL()

	//if err != nil {
	//	log.Panic("Problems with migrations", err)
	//	os.Exit(1)
	//}

	config.DB, err = pgx.Connect(context.Background(), pgString)
	//_,_ := config.DB.Ac
	//driver, err := postgres.WithInstance(config.DB, &postgres.Config{})
	//m, err := migrate.NewWithDatabaseInstance("file:///db/migrations", configDb.DBName, driver)
	//m.Steps(2)

	if err != nil {
		log.Panicln("Problems with connection to DB", err)
	}
	defer config.DB.Close(context.Background())

	r := routes.SetupRouter()

	r.Run()
}
