package studentOffice

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetViewKindnessScholarship 学生会查询应善良助学金
func GetViewKindnessScholarship(c *gin.Context) {
	id := c.Param("id")
	var studentOffice database.StudentOffice
	var ksarray []database.KindnessScholarship

	if dbe := database.DB.First(&studentOffice, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	if dbe := database.DB.Where("grade = ?", studentOffice.Grade).Find(&ksarray); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "ks": ksarray})
}

// PostViewKindnessScholarship 学生会更新应善良助学金状态
func PostViewKindnessScholarship(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	var ks database.KindnessScholarship

	if dbe := database.DB.First(&ks, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	ks.Status = status

	if dbe := database.DB.Model(&ks).Update(ks); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
