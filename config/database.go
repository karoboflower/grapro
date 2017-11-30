package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// DatabaseConfig 定义数据库用户名和密码
type DatabaseConfig struct {
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
}

// DBCfg 数据库配置的单例
var DBCfg *DatabaseConfig

// GetDBCfg 获取数据库配置实例
func GetDBCfg() *DatabaseConfig {
	if DBCfg == nil {
		DBCfg = &DatabaseConfig{}
		f, e := ioutil.ReadFile("config/database.json")
		if e != nil {
			fmt.Println("Open database.json failed!")
			os.Exit(1)
		}
		if e = json.Unmarshal(f, DBCfg); e != nil {
			fmt.Println("Parse database.json failed!")
			os.Exit(1)
		}
	}
	return DBCfg
}
