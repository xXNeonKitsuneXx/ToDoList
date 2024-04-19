package main

import (
	"fmt"

	"github.com/xXNeonKitsuneXx/toDoList_Backend/repository"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "BocchiKitsuNei:Crown1003@tcp(localhost:3306)/toDoListMariaDB?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	toDoListRepository := repository.NewToDoListRepositoryDB(db)

	// _ = toDoListRepository

	// todolists, err := toDoListRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(todolists)
	
	toDoListService := service.NewToDoListService(toDoListRepository)

	toDoLists, err := toDoListService.GetToDoLists()
	if err != nil {
		panic(err)
	}

	fmt.Println(toDoLists)
}
