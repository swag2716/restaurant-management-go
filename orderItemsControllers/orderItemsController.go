package orderItemsControllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/swapnika/restaurant-management/database"
	"github.com/swapnika/restaurant-management/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")
var validate = validator.New()
