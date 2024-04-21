package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/entities"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/handler"
	repository2 "github.com/xXNeonKitsuneXx/toDoList_Backend/internal/repository"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/service"
	"gorm.io/driver/mysql"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	initTimeZone()
	initConfig()
	//dsn := "BocchiKitsuNei:Crown1003@tcp(localhost:3306)/toDoListMariaDB?parseTime=true"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&entities.Todo{})
	if err != nil {
		panic("Failed to AutoMigrate")
	}

	toDoListRepositoryDB := repository2.NewToDoListRepositoryDB(db)
	_ = toDoListRepositoryDB

	// todolists, err := toDoListRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(todolists)

	toDoListRepositoryMock := repository2.NewToDoListRepositoryMock()
	_ = toDoListRepositoryMock
	toDoListService := service.NewToDoListService(toDoListRepositoryMock)

	toDoListHandler := handler.NewToDoListHandler(toDoListService)

	// toDoLists, err := toDoListService.GetToDoLists()
	// if err != nil {
	// 	panic(err)
	// }

	app := fiber.New()

	app.Get("/todolists", toDoListHandler.GetToDoLists)

	log.Printf("ToDoList run at port:  %v", viper.GetInt("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
