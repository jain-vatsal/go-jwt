package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-jwt/constants"
	"github.com/jain-vatsal/go-jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST(constants.UserSignUp, controllers.UserSignUp())
	incomingRoutes.POST(constants.UserLogin, controllers.UserLogin())
}
