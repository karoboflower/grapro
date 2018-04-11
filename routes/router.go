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
	router.LoadHTMLGlob("resources/views/*/*")
	router.GET("/", utilities.GetIndex)
	router.GET("/register", controller.GetRegister)
	router.POST("/register", controller.PostRegister)
	router.GET("/login", controller.GetLogin)
	router.POST("/login", controller.PostLogin)
	router.GET("/captcha/:filename", gin.WrapH(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
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
			studentsRouter.GET("/:id/SG", student.GetSG)
			studentsRouter.POST("/:id/SG", student.PostSG)
			studentsRouter.DELETE("/:id/SG", student.DeleteSG)
			// 国家励志奖学金
			studentsRouter.GET("/:id/NIS", student.GetNIS)
			studentsRouter.POST("/:id/NIS", student.PostNIS)
			studentsRouter.DELETE("/:id/NIS", student.DeleteNIS)
			// 应善良助学金
			studentsRouter.GET("/:id/KS", student.GetKS)
			studentsRouter.POST("/:id/KS", student.PostKS)
			studentsRouter.DELETE("/:id/KS", student.DeleteKS)
		}
		// 院学生会路由组
		studentOfficeRouter := authorized.Group("/2")
		{
			// 个人信息
			studentOfficeRouter.GET("/:id/profile", studentOffice.GetStudentOffice)
			studentOfficeRouter.POST("/:id/profile", studentOffice.PostStudentOffice)
			// 国家助学金
			studentOfficeRouter.GET("/:id/SG", studentOffice.GetSG)
			studentOfficeRouter.POST("/:id/SG", studentOffice.PostSG)
			// 应善良助学金
			studentOfficeRouter.GET("/:id/KS", studentOffice.GetKS)
			studentOfficeRouter.POST("/:id/KS", studentOffice.PostKS)
			// 国家励志奖学金
			studentOfficeRouter.GET("/:id/NIS", studentOffice.GetNIS)
			studentOfficeRouter.POST("/:id/NIS", studentOffice.PostNIS)
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
			counselorRouter.GET("/:id/SG", counselor.GetSG)
			counselorRouter.POST("/:id/SG", counselor.PostSG)
			// 国家励志奖学金
			counselorRouter.GET("/:id/NIS", counselor.GetNIS)
			counselorRouter.POST("/:id/NIS", counselor.PostNIS)
			// 应善良助学金
			counselorRouter.GET("/:id/KS", counselor.GetKS)
			counselorRouter.POST("/:id/KS", counselor.PostKS)
		}
	}
	return router
}
