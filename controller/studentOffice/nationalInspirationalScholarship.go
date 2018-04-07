package studentOffice

import (
	"gra-pro/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetViewNIS 学生会查询国家励志奖学金信息
func GetViewNIS(c *gin.Context) {
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

// PostViewNIS 学生会更新国家励志奖学金信息
func PostViewNIS(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	var nis database.NIS
	var i int
	var err error

	if dbe := database.DB.First(&nis, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	if i, err = strconv.Atoi(status); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	nis.Status = i

	if dbe := database.DB.Model(&nis).Update(nis); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
