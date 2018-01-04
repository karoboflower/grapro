package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Salt 密码盐
type Salt struct {
	Salt1 string `json:"salt1"`
	Salt2 string `json:"salt2"`
}

// GetSalt 获取密码盐
func GetSalt(InSalt *Salt) bool {

	f, err := ioutil.ReadFile("config/user/salt.json")

	if err != nil {
		log.Println("读取密码盐文件失败")
	}

	if err = json.Unmarshal(f, InSalt); err == nil {
		return true
	}

	return false
}
