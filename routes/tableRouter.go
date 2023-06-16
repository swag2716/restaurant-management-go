package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/tableControllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", tableControllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", tableControllers.GetTable())
	incomingRoutes.POST("/tables", tableControllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", tableControllers.UpdateTable())
}
