package utilities

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex 首页
func GetIndex(c *gin.Context) {
	var notifies []database.Notify
	database.DB.Find(&notifies)
	// c.HTML(http.StatusOK, "common/index.tmpl", gin.H{"notifies": notifies})
	c.HTML(http.StatusOK, "common/index.html", gin.H{"notifies": notifies})
}
