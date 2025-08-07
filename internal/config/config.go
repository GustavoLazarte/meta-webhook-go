package config

import (
	"log"
	"os"
	"webhook/internal/db"
	"webhook/internal/router"

	"github.com/joho/godotenv"
)

type Config struct {
	*router.Router
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	_, err = db.ConnectDatabase(dbHost)
	if err != nil {
		log.Fatal("Error when connecting to database")
	}
	config := &Config{
		Router: router.NewRouter(),
	}

	config.Router.SetupRouter()
	return config
}
