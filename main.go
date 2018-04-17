package main

import (
	"grapro/database"
	"grapro/routes"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	// Debug
	// log.Fatal(routes.Engine().Run(":8080"))
	// Release
	log.Fatal(routes.Engine().RunTLS(":8080", "./ca-certificates.crt", "./ca-certificates.key"))
}
