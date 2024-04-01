package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) CreateCustomers(c *gin.Context) {
	var requestBody dto.PostCustomersJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.CustomerRepo.CreateCustomers(c.Request.Context(), requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Received nil response"})
		return
	}

	for key, values := range response.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(response.StatusCode)
}

func (h Handler) GetCustomers(c *gin.Context) {
	var params dto.GetCustomersParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	response, err := h.CustomerRepo.GetCustomers(c.Request.Context(), &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Received nil response"})
		return
	}

	for key, values := range response.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(response.StatusCode)
}
