package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/restaurant-management/invoiceControllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", invoiceControllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", invoiceControllers.GetInvoice())
	incomingRoutes.POST("/invoices", invoiceControllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", invoiceControllers.UpdateInvoice())
}
