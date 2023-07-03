package orderItemsControllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/swapnika/restaurant-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")
var validate = validator.New()
