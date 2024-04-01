package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
)

type AuthHandler struct {
	AuthRepo adapters.AuthRepo
}

//
//func (h *AuthHandler) Login(c *gin.Context) {
//	var creds dto.Credentials
//	if err := c.ShouldBindJSON(&creds); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//		return
//	}
//
//	ctx := c.Request.Context()
//
//	token, err := h.AuthRepo.Login(ctx, creds.Username, creds.Password)
//	if err != nil {
//		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid username or password"})
//		return
//	}
//
//	c.JSON(http.StatusOK, token)
//}
//
//
//func NewAuthHandler(authRepo adapters.AuthRepo) *AuthHandler {
//	return &AuthHandler{
//		AuthRepo: authRepo,
//	}
//}
