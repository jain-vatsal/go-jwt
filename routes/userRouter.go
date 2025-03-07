package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/jain-vatsal/go-jwt/controllers"
	middleware "github.com/jain-vatsal/go-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthenticateUser())

	incomingRoutes.GET("/users", controllers.GetAllUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
