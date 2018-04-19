package studentOffice

import (
	"grapro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNIS 学生会查询国家励志奖学金信息
func GetNIS(c *gin.Context) {
	id := c.Param("id")
	var studentOffice database.StudentOffice
	var nisarray []database.NIS

	if dbe := database.DB.First(&studentOffice, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	if dbe := database.DB.Where("grade = ?", studentOffice.Grade).Find(&nisarray); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "nis": nisarray})
}

// PostNIS 学生会更新国家励志奖学金信息
func PostNIS(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	desc := c.PostForm("studentOfficeDesc")
	var nis database.NIS

	if dbe := database.DB.First(&nis, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	nis.Status = status
	nis.StudentOfficeDesc = desc

	if dbe := database.DB.Model(&nis).Update(nis); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "nis": nis})
}
