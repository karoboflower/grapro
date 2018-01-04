package routes

import (
	"gra-pro/controller"
	"gra-pro/database/models"

	"github.com/gin-gonic/gin"
)

// Engine Gin Engine
func Engine() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.POST("/api/register", controller.RegisterUser)
	router.POST("/api/login", models.UserPOST)
	return router
}
