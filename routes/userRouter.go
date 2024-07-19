package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-jwt/constants"
	"github.com/jain-vatsal/go-jwt/controllers"
	"github.com/jain-vatsal/go-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET(constants.GetAllUsers, controllers.GetAllUsers())
	incomingRoutes.GET(constants.GetUser, controllers.GetUser())

}
