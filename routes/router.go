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

	router := gin.Default()
	router.Use(favicon.New("./favicon.ico"))
	router.Delims("{%", "%}")
	router.LoadHTMLGlob("views/*/*")
	router.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "common/index.tmpl", gin.H{}) })
	router.GET("/register", controller.RegisterGET)
	router.POST("/register", controller.RegisterPOST)
	router.GET("/login", controller.LoginGET)
	router.POST("/login", controller.LoginPOST)
	authorized := router.Group("/auth")
	authorized.Use(middleware.JWTAuth(), middleware.RBAC())
	{
		authorized.GET("/student/:id/profile", controller.StudentGET)
		authorized.POST("/student/:id/profile", controller.StudentPOST)
		authorized.PUT("/student/:id/profile", controller.StudentPUT)
		authorized.GET("/student/:id/StateGrants", controller.StateGrantsGET)
		authorized.POST("/student/:id/StateGrants", controller.StateGrantsPOST)
		authorized.PUT("/student/:id/StateGrants", controller.StateGrantsPUT)
		authorized.DELETE("/student/:id/StateGrants", controller.StateGrantsDEL)
		authorized.GET("student/:id/NIS", controller.NISGET)
		authorized.POST("student/:id/NIS", controller.NISPOST)
		authorized.PUT("student/:id/NIS", controller.NISPUT)
		authorized.DELETE("student/:id/NIS", controller.NISDEL)
		authorized.GET("student/:id/KindnessScholarship", controller.KindnessScholarshipGET)
		authorized.POST("student/:id/KindnessScholarship", controller.KindnessScholarshipPOST)
		authorized.PUT("student/:id/KindnessScholarship", controller.KindnessScholarshipPUT)
		authorized.DELETE("student/:id/KindnessScholarship", controller.KindnessScholarshipDEL)
		authorized.GET("/counselor/:id/profile", controller.CounselorGET)
		authorized.POST("/counselor/:id/profile", controller.CounselorPOST)
		authorized.PUT("/counselor/:id/profile", controller.CounselorPUT)
		authorized.GET("/studentOffice/:id/profile", controller.StudentOfficeGET)
		authorized.POST("/studentOffice/:id/profile", controller.StudentOfficePOST)
		authorized.PUT("/studentOffice/:id/profile", controller.StudentOfficePUT)
		authorized.GET("/studentOffice/:id/notify", controller.NotifyGET)
		authorized.POST("/studentOffice/:id/notify", controller.NotifyPOST)
		authorized.PUT("/studentOffice/:id/notify", controller.NotifyPUT)
		authorized.DELETE("/studentOffice/:id/notify", controller.NotifyDEL)
	}
	return router
}
