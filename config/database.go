package config

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

// DB instance from gorm
var DB *pgx.Conn

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
	dbc.Host = os.Getenv("POSTGRES_HOST")
	dbc.Port = os.Getenv("POSTGRES_PORT")
	dbc.User = os.Getenv("POSTGRES_USER")
	dbc.Password = os.Getenv("POSTGRES_PASSWORD")
	dbc.DBName = os.Getenv("POSTGRES_DB")
}
