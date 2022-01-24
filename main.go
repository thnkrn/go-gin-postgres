package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	db "github.com/thnkrn/go-gin-postgres/config"
	auth "github.com/thnkrn/go-gin-postgres/middleware"
	routers "github.com/thnkrn/go-gin-postgres/routers"
)

func main() {
	router := gin.New()

	// Status Response
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO WORLD",
		})
	})

	// Request JWT
	router.POST("/login", auth.LoginHandler)

	database := db.Connect()

	// Authorization
	api := router.Group("/api", auth.AuthorizationMiddleware)
	routers.SetCollectionRoutes(api, database)

	fmt.Println("Server Running on Port: ", 3000)
	// http.ListenAndServe(":3000", router)
	router.Run(":3000")
}
