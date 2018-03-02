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
	router.Delims("{%", "%}")
	router.LoadHTMLGlob("views/*/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "common/index.tmpl", gin.H{
			"message": "huanglachuan",
		})
	})
	router.GET("/login", controller.LoginGET)
	router.POST("/login", controller.LoginPOST)
	router.GET("/register", controller.RegisterGET)
	router.POST("/register", controller.RegisterPOST)
	authorized := router.Group("/authorized")
	authorized.Use(middleware.JWTAuth())
	{
		authorized.GET("/:role/:id", controller.UserGET)
	}
	return router
}
