package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"social_system/service/controller"
)

func Init() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	Controller := new(controller.Controller)
	router.POST("/social_system/api", Controller.Handle)

	return router
}
