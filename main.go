package main

import (
	"github.com/gin-gonic/gin"

	db "github.com/thnkrn/go-gin-postgres/config"
	auth "github.com/thnkrn/go-gin-postgres/middleware"
	routers "github.com/thnkrn/go-gin-postgres/routers"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	// Status Response
	router.GET("/", func(c *gin.Context) {
		c.String(200, "ping pong")
	})

	// Request JWT
	router.POST("/login", auth.LoginHandler)

	database := db.Connect()

	// Authorization
	api := router.Group("/api", auth.AuthorizationMiddleware)
	routers.SetCollectionRoutes(api, database)

	// http.ListenAndServe(":3000", router)
	router.Run(":3000")
}
