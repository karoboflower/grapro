package main

import (
	"gra-pro/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	s := database.Student{
		ID: "2014051609033",
	}

	t := &database.Tester{
		ID:        2,
		StudentID: "2014051609033",
		Status:    1,
	}

	s.Delete(t)

	// Debug
	// log.Fatal(routes.Engine().Run(":8080"))
	// Release
	// log.Fatal(routes.Engine().RunTLS(":8080", "./ca-certificates.crt", "./ca-certificates.key"))
}
