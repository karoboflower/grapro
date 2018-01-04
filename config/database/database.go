// Package database 数据库相关配置信息
package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
