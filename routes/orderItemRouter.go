package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/orderItemsControllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", orderItemsControllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", orderItemsControllers.GetOrderItem())
	incomingRoutes.GET("orderItems-order/:order_id", orderItemsControllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", orderItemsControllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", orderItemsControllers.UpdateOrderItem())
}
