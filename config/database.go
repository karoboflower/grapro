package config

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
}

// UserData 数据库用户信息
type UserData struct {
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
}

// GetDBCfg 获取数据库配置信息及数据库用户信息,返回true则读取信息成功
func GetDBCfg(OutDBCfg *DBCfg, OutUserData *UserData) bool {
	// 读取数据库信息文件
	f, err := ioutil.ReadFile("config/databaseconfig.json")
	if err != nil {
		log.Fatal("读取数据库文件失败")
	}

	// 解析读取数据库信息
	if err = json.Unmarshal(f, OutDBCfg); err == nil {
		if err = json.Unmarshal(f, OutUserData); err == nil {
			return true
		}
	}
	return false
}
