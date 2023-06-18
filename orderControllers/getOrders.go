package orderControllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := orderCollection.Find(context.TODO(), bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
		}

		var allOrders []bson.M

		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, result)

	}
}
