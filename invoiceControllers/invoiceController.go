package invoiceControllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/swapnika/restaurant-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_no         interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var validate = validator.New()
