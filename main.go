package main

import (
	"fmt"
	"gra-pro/database"
	"reflect"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	s := database.Student{
		ID: "2014051609033",
	}

	t := &database.Tester{
		StudentID: "2014051609033",
		Status:    "0",
	}

	s.Create(t)
	if s.Read(t) {
		fmt.Println("success")
	} else {
		temp := reflect.ValueOf(t).Elem()
		tester := database.Tester{
			ID:        temp.Field(0).Int(),
			StudentID: temp.Field(1).String(),
			Status:    temp.Field(2).String(),
		}
		fmt.Println(tester.StudentID)
		fmt.Println(tester.Status)
	}

	// Debug
	// log.Fatal(routes.Engine().Run(":8080"))
	// Release
	// log.Fatal(routes.Engine().RunTLS(":8080", "./ca-certificates.crt", "./ca-certificates.key"))
}
