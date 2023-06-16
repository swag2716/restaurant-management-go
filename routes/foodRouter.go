package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/foodControllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", foodControllers.GetFoods())
	incomingRoutes.GET("/food/:food_id", foodControllers.GetFood())
	incomingRoutes.POST("/food", foodControllers.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", foodControllers.UpdateFood())
}
