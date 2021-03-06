package counselor

import (
	"grapro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNIS 辅导员查询国家励志奖学金信息
func GetNIS(c *gin.Context) {
	id := c.Param("id")
	var students []database.Student
	var temp []database.NIS
	var nisarray []database.NIS

	if dbe := database.DB.Where("counselor_id = ?", id).Find(&students); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	for _, student := range students {
		if dbe := database.DB.Where("student_id = ?", student.StudentID).Find(&temp); dbe != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
		for _, i := range temp {
			nisarray = append(nisarray, i)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "nis": nisarray})
}

// PostNIS 辅导员筛选国家励志奖学金信息
func PostNIS(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	desc := c.PostForm("counselorDesc")
	var nis database.NIS

	if dbe := database.DB.First(&nis, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	nis.Status = status
	nis.CounselorDesc = desc

	if dbe := database.DB.Model(&nis).Update(nis); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
