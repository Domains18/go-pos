package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) CreateSalesInvoices(c *gin.Context) {
	var requestBody dto.PostSalesInvoicesJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.InvoiceRepo.CreateSalesInvoices(c.Request.Context(), requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Received nil response"})
		return
	}

	// Set response headers
	for key, values := range response.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(response.StatusCode)
}
