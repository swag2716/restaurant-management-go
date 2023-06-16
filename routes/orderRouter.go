package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/orderControllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", orderControllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", orderControllers.GetOrder())
	incomingRoutes.POST("/orders", orderControllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", orderControllers.UpdateOrder())
}
