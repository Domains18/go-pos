package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (h Handler) GetProducts(c *gin.Context) {
	var params dto.GetProductsParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	response, err := h.ProductRepo.GetProducts(c.Request.Context(), &params)
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
	_, _ = io.Copy(c.Writer, response.Body)
}
