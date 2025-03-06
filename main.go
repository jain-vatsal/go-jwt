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
		port = constants.DEFAULTPORT
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for API-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for API-2"})
	})

	router.Run(":" + port)
}
