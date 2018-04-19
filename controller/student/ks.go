// Package student 应善良助学金
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

// GetKS 学生获取应善良助学金申请信息
func GetKS(c *gin.Context) {
	c.HTML(http.StatusOK, "student/ks.tmpl", gin.H{"id": c.Param("id")})
}

// PostKS 学生获提交善良助学金申请信息
func PostKS(c *gin.Context) {
	id := c.Param("id")
	var student database.Student
	var dst string
	var fileExtension string
	var ks database.KS

	if database.DB.First(&student, id).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "数据库没有您的个人信息"})
		return
	}
	// 应善良助学金材料存放目录
	dst = "resources/files/" + student.College + "/" + student.Profession + "/ks"
	if bExists, _ := utilities.PathExists(dst); !bExists {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			log.Fatal("Mkdir failed : " + err.Error())
		}
	}
	// 应善良助学金文件处理
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
		ks.Accessory = append(ks.Accessory, dst+id+strconv.Itoa(i)+"."+fileExtension)
		if err = c.SaveUploadedFile(file, dst+"/"+id+strconv.Itoa(i)+"."+fileExtension); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
			return
		}
	}
	ks.Status = "0"
	if dbe := database.DB.Create(&ks); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 0, "ks": ks})
}

// DeleteKS 学生删除应善良助学金申请信息
func DeleteKS(c *gin.Context) {
	id := c.PostForm("id")
	var ks database.KS

	database.DB.First(&ks, id)
	if dbe := database.DB.Where("ks_id = ?", id).Delete(&database.KS{}); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "ks": ks})
}
