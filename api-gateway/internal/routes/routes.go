package routes

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/routes/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.POST("/login", handlers.LoginCustomer(handler))
	router.POST("/customers", handler.CreateCustomers)
	router.GET("/customers", handler.GetCustomers)
	router.POST("/sales-invoices", handler.CreateSalesInvoices)
	router.GET("/products", handler.GetProducts)
}
