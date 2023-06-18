package orderControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var order models.Order
		var table models.Table

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if order.Table_id != nil {
			err := orderCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Table was not found"})
				return
			}

			updateObj = append(updateObj, bson.E{Key: "table", Value: order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})

		upsert := true
		filter := bson.M{"order_id": orderId}

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "order item update failed"})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}
