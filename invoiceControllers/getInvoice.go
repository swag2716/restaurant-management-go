package invoiceControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/models"
	"github.com/swapnika/restaurant-management/orderItemsControllers"
	"go.mongodb.org/mongo-driver/bson"
)

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice

		invoiceId := c.Param("invoice_id")
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var invoiceView InvoiceViewFormat

		allOrderItems, err := orderItemsControllers.ItemsByOrder(invoice.Order_id)

		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date
		invoiceView.Payment_status = invoice.Payment_status
		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Table_no = allOrderItems[0]["table_no"]
		invoiceView.Order_details = allOrderItems[0]["order_items"]

		c.JSON(http.StatusOK, invoiceView)

	}
}
