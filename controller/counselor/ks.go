package counselor

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKS 辅导员查询应善良助学金
func GetKS(c *gin.Context) {
	id := c.Param("id")
	var students []database.Student
	var temp database.KS
	var ksarray []database.KS

	if dbe := database.DB.Where("counselor_id = ?", id).Find(&students); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	for _, student := range students {
		if dbe := database.DB.First(&temp, student.StudentID); dbe != nil {
			continue
		}
		ksarray = append(ksarray, temp)
	}

	c.JSON(http.StatusOK, gin.H{"ks": ksarray})
}

// PostKS 辅导员筛选应善良助学金
func PostKS(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	var ks database.KS

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
