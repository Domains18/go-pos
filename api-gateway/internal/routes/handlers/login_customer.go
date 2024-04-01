package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func LoginCustomer(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody dto.PostLoginJSONRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		response, err := h.LoginRepo.LoginCustomer(c.Request.Context(), requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Read response body and headers
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		defer response.Body.Close()

		// Set response headers
		for key, values := range response.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}

		c.Status(response.StatusCode)
		c.Writer.Write(body)
	}
}
