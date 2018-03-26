package student

import (
	"fmt"
	"gra-pro/database"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetStateGrants 学生获取国家助学金申请信息
func GetStateGrants(c *gin.Context) {
	c.HTML(http.StatusOK, "student/stateGrants.tmpl", gin.H{})
}

// PostStateGrants 学生提交国家助学金申请信息
func PostStateGrants(c *gin.Context) {
	id := c.PostForm("id")
	// description := c.PostForm("description")
	var fileExtension string
	var student database.Student

	if database.DB.First(&student, id); student.ID != id {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "数据库没有您的信息"})
		return
	}

	fsquestionnaire, err := c.FormFile("fsquestionnaire")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if s := strings.Split(fsquestionnaire.Filename, "."); len(s) == 2 {
		fileExtension = s[1]
	}

	// accessory, err := c.FormFile("accessory")
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
	// 	return
	// }

	var dst string
	dst = "files/" + student.College + "/" + student.Profession
	if _, err = os.Stat(dst); err != nil {
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			fmt.Println("mkdir failed!", err.Error())
		}
	} else {
		fmt.Println("os.Stat report!", err.Error())
	}

	if err = c.SaveUploadedFile(fsquestionnaire, dst+"/"+id+"."+fileExtension); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 0})
}

// PutStateGrants 学生修改国家助学金申请信息
func PutStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// DeleteStateGrants 学生删除国家助学金申请信息
func DeleteStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
