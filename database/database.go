// Package database 数据库相关操作
package database

import (
	"gra-pro/config"
	"gra-pro/database/models"
	"log"

	"github.com/jinzhu/gorm"
)

// ConnectDB 连接数据库,返回*gorm.DB实例
func ConnectDB() *gorm.DB {
	// 数据库配置信息
	var dbcfg config.DBCfg
	// 数据库用户信息
	var userdata config.UserData

	if !config.GetDBCfg(&dbcfg, &userdata) {
		log.Fatalln("解析读取数据库信息失败")
	}

	db, e := gorm.Open(dbcfg.DBType, userdata.UserName+":"+userdata.Password+"@tcp("+dbcfg.IP+":"+dbcfg.Port+")/"+dbcfg.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if e != nil {
		log.Fatalln("数据库连接失败")
	}
	//if !db.HasTable(&models.User{}){
	db.AutoMigrate(&models.User{})
	//}
	return db
}
