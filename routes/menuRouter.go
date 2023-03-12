package routes

import (
	controller "Restaurant-Management/controllers"
	"Restaurant-Management/middleware"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menuss/:menu_id", controller.UpdateMenu())
}
