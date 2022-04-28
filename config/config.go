package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func LoadEnv(a *AppConfig) {
	err := godotenv.Load(".env")
	if err != nil {
		a.ErrorLog.Fatalln("Couldn't load the env file")
	}

	a.DbName = os.Getenv("DB_NAME")
	a.DbPort = os.Getenv("DB_PORT")
	a.DbUser = os.Getenv("DB_USER")
	a.DbPassword = os.Getenv("DB_PASSWORD")
	a.DbHost = os.Getenv("DB_HOST")

}
