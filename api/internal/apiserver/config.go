package apiserver

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPPort    string
	LogLevel    string
	DatabaseURL string
}

// NewConfig with default settings from the global environment variables of the env file
func NewConfig() *Config {
	prt := os.Getenv("API_PORT")
	lvl := os.Getenv("LOG_LEVEL")
	// connect to db
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return &Config{
		HTTPPort:    ":" + prt,
		LogLevel:    lvl,
		DatabaseURL: psqlconn,
	}
}
