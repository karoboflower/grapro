package main

import (
	"gra-pro/database"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	router.RunTLS(":8080", "./ca-certificates.crt", "./ca-certificates.key")
}
