package main

import (
	"github.com/Nerds-Catapult/notiwa/internal/server"
)

func main() {
	router := server.SetupRouter()
	router.Run(":8080")
}
