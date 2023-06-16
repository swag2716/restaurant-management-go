package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/menuControllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", menuControllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", menuControllers.GetMenu())
	incomingRoutes.POST("/menus", menuControllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", menuControllers.UpdateMenu())
}
