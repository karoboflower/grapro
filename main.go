package main

import (
	"gra-pro/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()
	//gin.SetMode(gin.DebugMode)

	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//"message": "Hello World",
	//})
	//})
	// listen and server on 0.0.0.0:8080
	//router.Run(":8080")
}
