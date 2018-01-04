// Package database 数据库相关操作
package database

import (
	"gra-pro/config/database"
	"gra-pro/database/models"
	"log"

	"github.com/jinzhu/gorm"
)

// DB *gorm.DB实例，全局变量
var DB *gorm.DB

// ConnectDB 连接数据库,返回*gorm.DB实例
func ConnectDB() {
	// 数据库配置信息
	var dbcfg database.DBCfg
	// 数据库用户信息
	var userdata database.UserData

	if !database.GetDBCfg(&dbcfg, &userdata) {
		log.Fatalln("解析读取数据库信息失败")
	}

	var e error

	DB, e = gorm.Open(dbcfg.DBType, userdata.UserName+":"+userdata.Password+"@tcp("+dbcfg.IP+":"+dbcfg.Port+")/"+dbcfg.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if e != nil {
		log.Fatalln("数据库连接失败")
	}
	//if !db.HasTable(&models.User{}){
	DB.AutoMigrate(&models.User{})
	//}
}
