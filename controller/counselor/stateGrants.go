package counselor

import (
	"gra-pro/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetViewStateGrants 辅导员查询国家助学金信息
func GetViewStateGrants(c *gin.Context) {
	id := c.Param("id")
	var students []database.Student
	var temp database.StateGrants
	var stateGrants []database.StateGrants

	if dbe := database.DB.Where("counselor_id = ?", id).Find(&students); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	for _, student := range students {
		if dbe := database.DB.First(&temp, student.ID); dbe != nil {
			continue
		}
		stateGrants = append(stateGrants, temp)
	}

	c.JSON(http.StatusOK, gin.H{"stateGrants": stateGrants})
}

// PostViewStateGrants 辅导员更改国家助学金状态
func PostViewStateGrants(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	var stateGrants database.StateGrants
	var i int
	var err error

	if dbe := database.DB.First(&stateGrants, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	if i, err = strconv.Atoi(status); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	stateGrants.Status = i

	if dbe := database.DB.Model(&stateGrants).Update(stateGrants); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
