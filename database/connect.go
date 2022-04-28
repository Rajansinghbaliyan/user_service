package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"user_service/config"
)

func ConnectDB(a *config.AppConfig) {

	port, err := strconv.ParseUint(a.DbPort, 10, 32)

	configData := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		a.DbHost, port, a.DbUser, a.DbPassword, a.DbName,
	)

	DB, err = gorm.Open(postgres.Open(configData), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to database")
	}

	a.InfoLog.Println("Connection opened to database")

}
