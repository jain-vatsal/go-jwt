package main

import (
	"os"

	"github.com/gin-gonic/gin"
	constants "github.com/jain-vatsal/go-jwt/constants"
	routes "github.com/jain-vatsal/go-jwt/routes"
)

func main() {
	port := os.Getenv(constants.PORT)
	if port == "" {
		port = constants.DEFAULT_PORT
	}

	router := gin.New()
	router.Use(gin.Logger())
	SetUpRouters(router, port)
}

func SetUpRouters(router *gin.Engine, port string) {
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
