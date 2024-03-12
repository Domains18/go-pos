package controller

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func HomeRoute(c *gin.Context, component templ.Component) error {
	c.Status(200)
	return component.Render(c.Request.Context(), c.Writer)
}
