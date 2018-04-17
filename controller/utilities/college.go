package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Specialty 专业信息
type Specialty struct {
	Name string   `json:"name"`
	Item []string `json:"item"`
}

// College 学院信息
type College struct {
	Sum         int         `json:"sum"`
	Specialties []Specialty `json:"college"`
}

// GetSpecialty 获取专业信息
func GetSpecialty(c *gin.Context) {
	var data College
	var college string

	college = c.PostForm("college")
	f, err := ioutil.ReadFile("resources/assets/student/json/college.json")
	if err != nil {
		log.Println("读取college.json失败")
	}
	if err = json.Unmarshal(f, &data); err == nil {
		for _, specialties := range data.Specialties {
			if college == specialties.Name {
				c.JSON(http.StatusOK, gin.H{"status": 0, "specialties": specialties.Item})
				return
			}
		}
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "读取专业信息失败"})
	}
}
