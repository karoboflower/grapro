// Package database 数据库的相关操作
package database

import (
	"gra-pro/config"
	"gra-pro/database/models"

	"github.com/jinzhu/gorm"
)

// ConnectDB 连接数据库,返回*gorm.DB实例
func ConnectDB(InDBCfg config.DBCfg, InUserData config.UserData) (*gorm.DB, error) {
	db, e := gorm.Open(InDBCfg.DBType, InUserData.UserName+":"+InUserData.Password+"@tcp("+InDBCfg.IP+":"+InDBCfg.Port+")/"+InDBCfg.DBName+"?charset=utf8&parseTime=True&loc=Local")
	//if !db.HasTable(&models.User{}){
	db.AutoMigrate(&models.User{})
	//}
	return db, e
}
