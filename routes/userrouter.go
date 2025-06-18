package routes

import (
	controllers "github.com/ramkrishnareddy24/GoRestaurant/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.SignUo())
	incomingRoutes.POST("/users/login", controllers.Login())

}