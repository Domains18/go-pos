package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	AuthRepo adapters.AuthRepo
}


func (h *Handler) Login(c *gin.Context) {
	var creds dto.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := c.Request.Context()

	token, err := h.AuthRepo.Login(ctx, creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, token)
}


func NewHandler(authRepo adapters.AuthRepo) *Handler {
	return &Handler{
		AuthRepo: authRepo,
	}
}