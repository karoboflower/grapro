package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Secret JWT签署密钥
type Secret struct {
	SignKey string `json:"secret"`
}

// GetSignKey 获取JWT签署密钥
func GetSignKey() (Secret, bool) {
	var secret Secret

	f, err := ioutil.ReadFile("config/auth/secret.json")
	if err != nil {
		log.Fatalln("读取JWT密钥文件失败")
	}
	if err = json.Unmarshal(f, &secret); err == nil {
		return secret, true
	}
	return secret, false
}
