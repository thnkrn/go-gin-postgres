package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	routers "github.com/thnkrn/go-gin-postgres/routers"
)

func main() {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	api := router.Group("/api")
	routers.SetCollectionRoutes(api, db)

	fmt.Println("Server Running on Port: ", 3000)
	http.ListenAndServe(":3000", router)
}
