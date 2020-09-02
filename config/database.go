package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// DB instance from gorm
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// DbURL returns string db connection
func (dbc *DBConfig) DbURL() string {
	return fmt.Sprintf(
		// "postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		dbc.User,
		dbc.Password,
		dbc.DBName,
		dbc.Port,
		dbc.Host,
	)
}

// BuildDBConfig is used for building db config
func (dbc *DBConfig) BuildDBConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	dbc.Host = host
	dbc.Port = port
	dbc.User = user
	dbc.Password = password
	dbc.DBName = dbName
}
