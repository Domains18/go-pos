package controller

import (
	"github.com/gin-gonic/gin"
)

// TODO: configure to return templ files and hydated HTML file instead of JSON responses
func HomeRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}
