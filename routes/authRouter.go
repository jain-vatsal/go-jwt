package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/jain-vatsal/go-jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("users/sign-up", controllers.SignUpUser())
	incomingRoutes.POST("users/login", controllers.LoginUser())

}
