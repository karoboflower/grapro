package routes

import (
	"gra-pro/controller"
	"gra-pro/controller/counselor"
	"gra-pro/controller/student"
	"gra-pro/controller/studentOffice"
	"gra-pro/controller/utilities"
	"gra-pro/middleware"

	"net/http"

	"github.com/dchest/captcha"
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
	router.GET("/", utilities.GetIndex)
	router.GET("/register", controller.GetRegister)
	router.POST("/register", controller.PostRegister)
	router.GET("/login", controller.GetLogin)
	router.POST("/login", controller.PostLogin)
	router.GET("/captcha/:fileName", gin.WrapH(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	router.POST("/captcha", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"captcha": captcha.New()}) })
	router.GET("/notify", utilities.GetNotify)
	authorized := router.Group("/auth")
	authorized.Use(middleware.JWTAuth(), middleware.RBAC())
	{
		// 校学生会路由组
		schoolRouter := authorized.Group("/4")
		{
			schoolRouter.GET("/:id/profile", func(c *gin.Context) { c.JSON(http.StatusOK, "hello") })
		}
		// 学生路由组
		studentsRouter := authorized.Group("/3")
		{
			// 个人信息
			studentsRouter.GET("/:id/profile", student.GetProfile)
			studentsRouter.POST("/:id/profile", student.PostProfile)
			studentsRouter.GET("/:id/GetCounselor", student.GetCounselor)
			// 国家助学金
			studentsRouter.GET("/:id/StateGrants", student.GetStateGrants)
			studentsRouter.POST("/:id/StateGrants", student.PostStateGrants)
			studentsRouter.DELETE("/:id/StateGrants", student.DeleteStateGrants)
			// 国家励志奖学金
			studentsRouter.GET("/:id/NIS", student.GetNIS)
			studentsRouter.POST("/:id/NIS", student.PostNIS)
			studentsRouter.DELETE("/:id/NIS", student.DeleteNIS)
			// 应善良助学金
			studentsRouter.GET("/:id/KindnessScholarship", student.GetKindnessScholarship)
			studentsRouter.POST("/:id/KindnessScholarship", student.PostKindnessScholarship)
			studentsRouter.DELETE("/:id/KindnessScholarship", student.DeleteKindnessScholarship)
		}
		// 院学生会路由组
		studentOfficeRouter := authorized.Group("/2")
		{
			// 个人信息
			studentOfficeRouter.GET("/:id/profile", studentOffice.GetStudentOffice)
			studentOfficeRouter.POST("/:id/profile", studentOffice.PostStudentOffice)
			// 国家助学金
			studentOfficeRouter.GET("/:id/ViewStateGrants", studentOffice.GetViewStateGrants)
			studentOfficeRouter.POST("/:id/ViewStateGrants", studentOffice.PostViewStateGrants)
			// 应善良助学金
			studentOfficeRouter.GET("/:id/ViewKindnessScholarship", studentOffice.GetViewKindnessScholarship)
			studentOfficeRouter.POST("/:id/ViewKindnessScholarship", studentOffice.PostViewKindnessScholarship)
			// 国家励志奖学金
			studentOfficeRouter.GET("/:id/ViewNIS", studentOffice.GetViewNIS)
			studentOfficeRouter.POST("/:id/ViewNIS", studentOffice.PostViewNIS)
			// 通知
			studentOfficeRouter.POST("/:id/notify", studentOffice.PostNotify)
			studentOfficeRouter.DELETE("/:id/notify", studentOffice.DeleteNotify)
		}
		// 辅导员路由组
		counselorRouter := authorized.Group("/1")
		{
			// 个人信息
			counselorRouter.GET("/:id/profile", counselor.GetCounselor)
			counselorRouter.POST("/:id/profile", counselor.PostCounselor)
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
