// Package database 数据库相关操作
package database

import (
	"gra-pro/models"

	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
)

// DBCfg 数据库配置信息
type DBCfg struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	DBType string `json:"db_type"`
	DBName string `json:"db_name"`
}

// UserData 数据库用户信息
type UserData struct {
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
}

// DB *gorm.DB实例，全局变量
var DB *gorm.DB

// GetDBCfg 获取数据库配置信息及数据库用户信息
func GetDBCfg(InDBCfg *DBCfg, InUserData *UserData) bool {
	// 读取数据库信息文件
	f, err := ioutil.ReadFile("config/database/databaseconfig.json")
	if err != nil {
		log.Fatalln("读取数据库文件失败")
	}

	// 解析读取数据库信息
	if err = json.Unmarshal(f, InDBCfg); err == nil {
		if err = json.Unmarshal(f, InUserData); err == nil {
			return true
		}
	}
	return false
}

// ConnectDB 连接数据库,返回*gorm.DB实例
func ConnectDB() {
	// 数据库配置信息
	var dbcfg DBCfg
	// 数据库用户信息
	var userdata UserData

	if !GetDBCfg(&dbcfg, &userdata) {
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