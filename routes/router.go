package routes

import (
	"gra-pro/controller"
	"gra-pro/controller/counselor"
	"gra-pro/controller/student"
	"gra-pro/controller/studentOffice"
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
		authorized.GET("/student/:id/profile", student.GetProfile)
		authorized.POST("/student/:id/profile", student.PostProfile)
		authorized.PUT("/student/:id/profile", student.PutProfile)
		authorized.GET("/student/:id/StateGrants", student.GetStateGrants)
		authorized.POST("/student/:id/StateGrants", student.PostStateGrants)
		authorized.PUT("/student/:id/StateGrants", student.PutStateGrants)
		authorized.DELETE("/student/:id/StateGrants", student.DeleteStateGrants)
		authorized.GET("student/:id/NIS", student.GetNIS)
		authorized.POST("student/:id/NIS", student.PostNIS)
		authorized.PUT("student/:id/NIS", student.PutNIS)
		authorized.DELETE("student/:id/NIS", student.DeleteNIS)
		authorized.GET("student/:id/KindnessScholarship", student.GetKindnessScholarship)
		authorized.POST("student/:id/KindnessScholarship", student.PostKindnessScholarship)
		authorized.PUT("student/:id/KindnessScholarship", student.PutKindnessScholarship)
		authorized.DELETE("student/:id/KindnessScholarship", student.DeleteKindnessScholarship)
		authorized.GET("/counselor/:id/profile", counselor.GetCounselor)
		authorized.POST("/counselor/:id/profile", counselor.PostCounselor)
		authorized.PUT("/counselor/:id/profile", counselor.PutCounselor)
		authorized.GET("counselor/:id/ViewStateGrants", counselor.GetStateGrants)
		authorized.POST("counselor/:id/ViewStateGrants", counselor.PostStateGrants)
		authorized.GET("counselor/:id/ViewNIS", counselor.GetNIS)
		authorized.POST("counselor/:id/ViewNIS", counselor.PostNIS)
		authorized.GET("counselor/:id/ViewKindnessScholarship", counselor.GetKindnessScholarship)
		authorized.POST("counselor/:id/ViewKindnessScholarship", counselor.PostKindnessScholarship)
		authorized.GET("/studentOffice/:id/profile", studentOffice.GetStudentOffice)
		authorized.POST("/studentOffice/:id/profile", studentOffice.PostStudentOffice)
		authorized.PUT("/studentOffice/:id/profile", studentOffice.PutStudentOffice)
		authorized.GET("/studentOffice/:id/notify", studentOffice.GetNotify)
		authorized.POST("/studentOffice/:id/notify", studentOffice.PostNotify)
		authorized.PUT("/studentOffice/:id/notify", studentOffice.PutNotify)
		authorized.DELETE("/studentOffice/:id/notify", studentOffice.DeleteNotify)
	}
	return router
}
