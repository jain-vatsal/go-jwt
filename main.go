package main

import (
	"os"

	"github.com/gin-gonic/gin"
	cns "github.com/jain-vatsal/go-jwt/constants"
	routes "github.com/jain-vatsal/go-jwt/routes"
)

func main() {
	port := os.Getenv(cns.PORT)
	if port == "" {
		port = cns.DEFAULTPORT
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET(cns.API_1, func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for API-1"})
	})

	router.GET(cns.API_1, func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for API-2"})
	})

	router.Run(":" + port)
}
