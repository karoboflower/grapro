package counselor

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetViewNIS 辅导员查询国家励志奖学金信息
func GetViewNIS(c *gin.Context) {
	id := c.Param("id")
	var students []database.Student
	var temp database.NIS
	var nisarray []database.NIS

	if dbe := database.DB.Where("counselor_id = ?", id).Find(&students); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	for _, student := range students {
		if dbe := database.DB.First(&temp, student.StudentID); dbe != nil {
			continue
		}
		nisarray = append(nisarray, temp)
	}

	c.JSON(http.StatusOK, gin.H{"nis": nisarray})
}

// PostViewNIS 辅导员筛选国家励志奖学金信息
func PostViewNIS(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	var nis database.NIS

	if dbe := database.DB.First(&nis, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	nis.Status = status

	if dbe := database.DB.Model(&nis).Update(nis); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
