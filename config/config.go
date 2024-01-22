package config

import (
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)


type (
	Config struct {
		App App
		// Db  Db
		// Jwt Jwt
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file : %s", err.Error())
	}

	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
	}

}

