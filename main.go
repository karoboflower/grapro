package main

import (
	"gra-pro/config"
	"gra-pro/database"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 数据库用户信息实例
	var userdata config.UserData
	// 数据库配置信息实例
	var dbcfg config.DBCfg

	if !config.GetDBCfg(&dbcfg, &userdata) {
		log.Fatalln("解析读取数据库信息失败")
	}

	db, err := database.ConnectDB(dbcfg, userdata)
	if err != nil {
		log.Fatalln("数据库连接失败")
	}
	defer db.Close()
	//gin.SetMode(gin.DebugMode)

	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//"message": "Hello World",
	//})
	//})
	// listen and server on 0.0.0.0:8080
	//router.Run(":8080")
}
