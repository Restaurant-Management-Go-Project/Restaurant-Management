package routes

import (
	controller "Restaurant-Management/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/foods", controller.GetFoods())
}
