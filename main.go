package main

import (
	"fmt"
	"gra-pro/config"
)

func main() {
	fmt.Println(config.GetDBCfg().Password)
	//gin.SetMode(gin.DebugMode)

	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//"message": "Hello World",
	//})
	//})
	//router.Run(":8080") // listen and server on 0.0.0.0:8080
}
