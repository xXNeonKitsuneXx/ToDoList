package main

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/entities"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/handler"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/repository"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "BocchiKitsuNei:Crown1003@tcp(localhost:3306)/toDoListMariaDB?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&entities.Todo{})
	if err != nil {
		panic("Failed to AutoMigrate")
	}

	toDoListRepositoryDB := repository.NewToDoListRepositoryDB(db)
	_ = toDoListRepositoryDB

	// todolists, err := toDoListRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(todolists)

	toDoListRepositoryMock := repository.NewToDoListRepositoryMock()
	toDoListService := service.NewToDoListService(toDoListRepositoryMock)

	toDoListHandler := handler.NewToDoListHandler(toDoListService)

	// toDoLists, err := toDoListService.GetToDoLists()
	// if err != nil {
	// 	panic(err)
	// }

	app := fiber.New()

	app.Get("/todolists", toDoListHandler.GetToDoLists)

	app.Listen(":8000")

}
