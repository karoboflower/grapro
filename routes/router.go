package routes

import (
	"gra-pro/controller"
	"gra-pro/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// Engine Gin Engine
func Engine() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(favicon.New("./favicon.ico"))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.tmpl", gin.H{
			"title": "huanglachuan",
		})
	})
	router.POST("/api/register", controller.RegisterUser)
	router.POST("/api/login", controller.UserPOST)
	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.POST("/api/user/:id", controller.UserGET)
	}
	return router
}
