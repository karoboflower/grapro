package studentOffice

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSG 学生会查询国家助学金信息
func GetSG(c *gin.Context) {
	id := c.Param("id")
	var studentOffice database.StudentOffice
	var SG []database.SG

	if dbe := database.DB.First(&studentOffice, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	if dbe := database.DB.Where("grade = ?", studentOffice.Grade).Find(&SG); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "SG": SG})
}

// PostSG 学生会更改国家助学金状态
func PostSG(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	var SG database.SG

	if dbe := database.DB.First(&SG, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	SG.Status = status

	if dbe := database.DB.Model(&SG).Update(SG); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
