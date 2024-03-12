package server

import (

	// "github.com/Nerds-Catapult/notiwa/view"
	"github.com/Nerds-Catapult/notiwa/internal/server/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		controller.HomeRoute(c, view.Home.HomePage)
	})
	return router
}
