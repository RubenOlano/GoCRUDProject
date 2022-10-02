package routes

import (
	"CRUD/controller"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine) {
	router.POST("/movies", controller.CreateMovie())
	router.GET("/movies", controller.GetMovie())
	router.PUT("/movies/:id", controller.EditMovie())
	router.DELETE("/movies/:id", controller.DeleteMovie())
}
