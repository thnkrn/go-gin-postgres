package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/thnkrn/go-gin-postgres/config"
	routers "github.com/thnkrn/go-gin-postgres/routers"
)

func main() {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})

	database := db.Connect()

	api := router.Group("/api")
	routers.SetCollectionRoutes(api, database)

	fmt.Println("Server Running on Port: ", 3000)
	http.ListenAndServe(":3000", router)
}
