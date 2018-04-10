package main

import (
	"fmt"
	"gra-pro/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	s := database.Student{
		StudentID: "2014051609034",
	}

	var t []database.Tester

	if s.ReadAll(&t) {
		fmt.Println("success")
		for _, tester := range t {
			fmt.Println(tester.ID)
		}
	}
	fmt.Println("failed")
	// if dbe := database.DB.Where("student_id = ?", s.ID).Find(&t); dbe.Error != nil {
	// 	fmt.Println(dbe.Error)
	// }
	// fmt.Println("success")

	// Debug
	// log.Fatal(routes.Engine().Run(":8080"))
	// Release
	// log.Fatal(routes.Engine().RunTLS(":8080", "./ca-certificates.crt", "./ca-certificates.key"))
}
