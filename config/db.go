package config

import (
	"fmt"
	"log"

	models "github.com/thnkrn/go-gin-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	config, err := LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		fmt.Println(dbErr.Error())
		panic("Cannot connect to Database")
	}

	db.AutoMigrate(&models.Campaigns{})

	return db
}
