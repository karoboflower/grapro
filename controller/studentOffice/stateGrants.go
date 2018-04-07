package studentOffice

import (
	"gra-pro/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetViewStateGrants 学生会查询国家助学金信息
func GetViewStateGrants(c *gin.Context) {
	id := c.Param("id")
	var studentOffice database.StudentOffice
	var stateGrants []database.StateGrants

	if dbe := database.DB.First(&studentOffice, id); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	if dbe := database.DB.Where("grade = ?", studentOffice.Grade).Find(&stateGrants); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "stateGrants": stateGrants})
}

// PostViewStateGrants 学生会更改国家助学金状态
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
