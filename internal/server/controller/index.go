package controller

import (
	"github.com/gin-gonic/gin"
)

func HomeRoute(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "world"})
}
