package studentOffice

import (
	"grapro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKS 学生会查询应善良助学金
func GetKS(c *gin.Context) {
	id := c.Param("id")
	var studentOffice database.StudentOffice
	var ksarray []database.KS

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

// PostKS 学生会更新应善良助学金状态
func PostKS(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	desc := c.PostForm("studentOfficeDesc")
	var ks database.KS

	if dbe := database.DB.First(&ks, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	ks.Status = status
	ks.StudentOfficeDesc = desc

	if dbe := database.DB.Model(&ks).Update(ks); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "ks": ks})
}
