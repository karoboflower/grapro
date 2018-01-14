package main

import (
	"gra-pro/database"
	"gra-pro/routes"

	"log"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	a := gormadapter.NewAdapter("mysql", "root:DelusionHost520@tcp(127.0.0.1:3306)/grapro", true)
	authEnforcer, err := casbin.NewEnforcerSafe("./config/auth/auth_model.conf", a)
	if err != nil {
		log.Fatalln(err)
	}
	authEnforcer.LoadPolicy()
	authEnforcer.SavePolicy()
	authEnforcer.AddPolicy()

	// Debug
	log.Fatal(routes.Engine().Run(":8080"))
	// Release
	// log.Fatal(routes.Engine().RunTLS(":8080", "./ca-certificates.crt", "./ca-certificates.key"))
}
