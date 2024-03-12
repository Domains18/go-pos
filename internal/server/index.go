package server

import (
	"github.com/Nerds-Catapult/notiwa/internal/server/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.HomeRoute)

	return router
}
