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
	router.MaxMultipartMemory = 8 << 20
	router.Use(favicon.New("./favicon.ico"))
	router.Delims("{%", "%}")
	router.LoadHTMLGlob("views/*/*")
	router.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "common/index.tmpl", gin.H{}) })
	router.GET("/register", controller.GetRegister)
	router.POST("/register", controller.PostRegister)
	router.GET("/login", controller.GetLogin)
	router.POST("/login", controller.PostLogin)
	authorized := router.Group("/auth")
	authorized.Use(middleware.JWTAuth(), middleware.RBAC())
	{
		// 学生路由组
		studentsRouter := authorized.Group("/3")
		{
			// 个人信息
			studentsRouter.GET("/:id/profile", student.GetProfile)
			studentsRouter.POST("/:id/profile", student.PostProfile)
			studentsRouter.PUT("/:id/profile", student.PutProfile)
			// 国家助学金
			studentsRouter.GET("/:id/StateGrants", student.GetStateGrants)
			studentsRouter.POST("/:id/StateGrants", student.PostStateGrants)
			studentsRouter.PUT("/:id/StateGrants", student.PutStateGrants)
			studentsRouter.DELETE("/:id/StateGrants", student.DeleteStateGrants)
			// 国家励志奖学金
			studentsRouter.GET("/:id/NIS", student.GetNIS)
			studentsRouter.POST("/:id/NIS", student.PostNIS)
			studentsRouter.PUT("/:id/NIS", student.PutNIS)
			studentsRouter.DELETE("/:id/NIS", student.DeleteNIS)
			// 应善良助学金
			studentsRouter.GET("/:id/KindnessScholarship", student.GetKindnessScholarship)
			studentsRouter.POST("/:id/KindnessScholarship", student.PostKindnessScholarship)
			studentsRouter.PUT("/:id/KindnessScholarship", student.PutKindnessScholarship)
			studentsRouter.DELETE("/:id/KindnessScholarship", student.DeleteKindnessScholarship)
		}
		// 学生会路由组
		studentOfficeRouter := authorized.Group("/2")
		{
			// 个人信息
			studentOfficeRouter.GET("/:id/profile", studentOffice.GetStudentOffice)
			studentOfficeRouter.POST("/:id/profile", studentOffice.PostStudentOffice)
			studentOfficeRouter.PUT("/:id/profile", studentOffice.PutStudentOffice)
			// 通知
			studentOfficeRouter.GET("/:id/notify", studentOffice.GetNotify)
			studentOfficeRouter.POST("/:id/notify", studentOffice.PostNotify)
			studentOfficeRouter.PUT("/:id/notify", studentOffice.PutNotify)
			studentOfficeRouter.DELETE("/:id/notify", studentOffice.DeleteNotify)
		}
		// 辅导员路由组
		counselorRouter := authorized.Group("/1")
		{
			// 个人信息
			counselorRouter.GET("/:id/profile", counselor.GetCounselor)
			counselorRouter.POST("/:id/profile", counselor.PostCounselor)
			counselorRouter.PUT("/:id/profile", counselor.PutCounselor)
			// 国家助学金
			counselorRouter.GET("/:id/ViewStateGrants", counselor.GetViewStateGrants)
			counselorRouter.POST("/:id/ViewStateGrants", counselor.PostViewStateGrants)
			// 国家励志奖学金
			counselorRouter.GET("/:id/ViewNIS", counselor.GetViewNIS)
			counselorRouter.POST("/:id/ViewNIS", counselor.PostViewNIS)
			// 应善良助学金
			counselorRouter.GET("/:id/ViewKindnessScholarship", counselor.GetViewKindnessScholarship)
			counselorRouter.POST("/:id/ViewKindnessScholarship", counselor.PostViewKindnessScholarship)
		}
	}
	return router
}
