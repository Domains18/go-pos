package cmd

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/repository"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/routes"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/routes/handlers"
	"github.com/Nerds-Catapult/notiwa/api-gateway/pkg/db/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Execute() {
	postgres.InitDB()

	// Initialize repositories
	authRepo := repository.NewAuthRepo(postgres.DB)
	customerRepo := repository.NewCustomerRepo(postgres.DB)
	invoiceRepo := repository.NewInvoiceRepo(postgres.DB)
	loginRepo := repository.NewLoginRepo(postgres.DB)

	// Initialize handlers
	handler := &handlers.Handler{
		AuthRepo:     authRepo,
		CustomerRepo: customerRepo,
		InvoiceRepo:  invoiceRepo,
		LoginRepo:    loginRepo,
	}

	router := gin.Default()

	routes.SetupRoutes(router, handler)

	// Start server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
