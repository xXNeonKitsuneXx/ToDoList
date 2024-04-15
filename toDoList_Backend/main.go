package main

import (
	"fmt"

	"github.com/xXNeonKitsuneXx/toDoList_Backend/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// import "fmt"

// func main() {
// 	fmt.Println("PEKOPEKO")
// 	fmt.Println(Hello("Trainer-San"));
// }

// func Hello(name string) string {
//     message := fmt.Sprintf("Konbanwa, %v. Irachai!", name)
//     return message
// }

func main() {
	dsn := "BocchiKitsuNei:Crown1003@tcp(localhost:3306)/toDoListMariaDB?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	toDoListRepository := repository.NewToDoListRepositoryDB(db)

	_ = toDoListRepository

	todolists, err := toDoListRepository.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(todolists)
}
