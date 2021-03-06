// Package student 国家励志奖学金
package student

import (
	"grapro/database"
	"grapro/utilities"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetNIS 学生获取国家励志奖学金申请信息
func GetNIS(c *gin.Context) {
	var nis []database.NIS
	if dbe := database.DB.Where("student_id = ?", c.Param("id")).Find(&nis); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}
	c.HTML(http.StatusOK, "student/nis.tmpl", gin.H{"status": 0, "id": c.Param("id"), "nis": nis})
}

// PostNIS 学生提交国家励志奖学金申请信息
func PostNIS(c *gin.Context) {
	id := c.Param("id")
	var student database.Student
	var dst string
	var fileExtension string
	var nis database.NIS

	if database.DB.First(&student, id).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "数据库没有您的个人信息"})
		return
	}
	// 国家励志奖学金材料存放目录
	dst = "resources/files/" + student.College + "/" + student.Profession + "/NIS"
	if bExists, _ := utilities.PathExists(dst); !bExists {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			log.Fatal("Mkdir failed : " + err.Error())
		}
	}
	// 国家励志奖学金文件处理
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	accessory := form.File["accessory"]
	for i, file := range accessory {
		if s := strings.Split(file.Filename, "."); len(s) == 2 {
			fileExtension = s[1]
		}
		nis.Accessory = append(nis.Accessory, dst+id+strconv.Itoa(i)+"."+fileExtension)
		if err = c.SaveUploadedFile(file, dst+"/"+id+strconv.Itoa(i)+"."+fileExtension); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
			return
		}
	}
	nis.Status = "0"
	if dbe := database.DB.Create(&nis); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 0, "nis": nis})
}

// DeleteNIS 学生删除国家励志奖学金申请信息
func DeleteNIS(c *gin.Context) {
	id := c.PostForm("id")
	var nis database.NIS

	database.DB.First(&nis, id)
	if dbe := database.DB.Where("nis_id = ?", id).Delete(&database.NIS{}); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "nis": nis})
}
