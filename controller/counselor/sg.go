package counselor

import (
	"grapro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSG 辅导员查询国家助学金信息
func GetSG(c *gin.Context) {
	id := c.Param("id")
	var students []database.Student
	var temp database.SG
	var SG []database.SG

	if dbe := database.DB.Where("counselor_id = ?", id).Find(&students); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	for _, student := range students {
		if dbe := database.DB.First(&temp, student.StudentID); dbe != nil {
			continue
		}
		SG = append(SG, temp)
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "SG": SG})
}

// PostSG 辅导员更改国家助学金状态
func PostSG(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	desc := c.PostForm("counselorDesc")
	var SG database.SG

	if dbe := database.DB.First(&SG, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	SG.Status = status
	SG.CounselorDesc = desc

	if dbe := database.DB.Model(&SG).Update(SG); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "SG": SG})
}
