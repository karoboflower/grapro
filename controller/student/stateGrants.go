package student

import (
	"gra-pro/database"
	"gra-pro/utilities"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetStateGrants 学生获取国家助学金申请信息
func GetStateGrants(c *gin.Context) {
	c.HTML(http.StatusOK, "student/stateGrants.tmpl", gin.H{"id": c.Param("id")})
}

// PostStateGrants 学生提交国家助学金申请信息
func PostStateGrants(c *gin.Context) {
	id := c.PostForm("id")
	var student database.Student
	var fileExtension string
	var dst string
	var stateGrants database.StateGrants

	if database.DB.First(&student, id); student.ID != id {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "数据库没有您的个人信息"})
		return
	}

	stateGrants.StudentID = id
	fsquestionnaire, err := c.FormFile("fsquestionnaire")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if s := strings.Split(fsquestionnaire.Filename, "."); len(s) == 2 {
		fileExtension = s[1]
	}

	// 申请助学金必须材料家庭情况调查表存放目录
	dst = "files/" + student.College + "/" + student.Profession + "/fsq"
	if bExists, _ := utilities.PathExists(dst); !bExists {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			log.Fatal("Mkdir failed : " + err.Error())
		}
	}

	stateGrants.FSQuestionnaire = id + "." + fileExtension
	if err = c.SaveUploadedFile(fsquestionnaire, dst+"/"+id+"."+fileExtension); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	// 处理助学金申请附件
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	// 助学金材料附件存放目录
	dst = "files/" + student.College + "/" + student.Profession + "/accessory"
	if bExists, _ := utilities.PathExists(dst); !bExists {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			log.Fatal("Mkdir failed : " + err.Error())
		}
	}
	accessory := form.File["accessory"]
	for i, file := range accessory {
		if s := strings.Split(file.Filename, "."); len(s) == 2 {
			fileExtension = s[1]
		}
		stateGrants.Accessory = stateGrants.Accessory + id + strconv.Itoa(i) + "." + fileExtension + ";"
		if err = c.SaveUploadedFile(fsquestionnaire, dst+"/"+id+strconv.Itoa(i)+"."+fileExtension); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
			return
		}
	}

	stateGrants.Status = 0
	if dbe := database.DB.Create(&stateGrants); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
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
