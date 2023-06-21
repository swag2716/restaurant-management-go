package orderControllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/swapnika/restaurant-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")
var validate = validator.New()
