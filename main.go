package main

import (
	"CRUD/config"

	"github.com/gin-gonic/gin"

	"CRUD/routes"
)

func main() {
	router := gin.Default()

	routes.MovieRoute(router)

	config.ConnectDB()

	router.Run(":8080")
}
