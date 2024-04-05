// package main

// import "fmt"

// func main() {
// 	fmt.Println("PEKOPEKO")
// 	fmt.Println(Hello("Trainer-San"));
// }

// func Hello(name string) string {
//     message := fmt.Sprintf("Konbanwa, %v. Irachai!", name)
//     return message
// }

////////////////////////////////////////////////////////////

// package main

// import "github.com/gofiber/fiber/v2"

// func main() {
//     app := fiber.New()

//     app.Get("/", func(c *fiber.Ctx) error {
//         return c.SendString("KONKONKITSUNE!")
//     })

//     app.Listen(":3000")
// }

//////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type ListNeedToDo struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DoneAt      *time.Time `gorm:"null"`
	Name        string
	DeadlineAt  time.Time
	Description string
	DoneInTime  int
}

func main() {
	dsn := "postgresql://bocchikitsunei:C2LUJGWd5wn8xJDEmvo6Ng@todolistbocchikitsunei-9025.8nk.gcp-asia-southeast1.cockroachlabs.cloud:26257/ToDoList_BocchiKitsuNei?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	fmt.Println(now)

	// Migrate the schema
	err = db.AutoMigrate(&ListNeedToDo{})
	if err != nil {
		return
	}

	// Create
	db.Create(&ListNeedToDo{Name: "Name001", CreatedAt: time.Now(), DeadlineAt: time.Now().Add(24 * time.Hour), Description: "This is an example task."})

	//// Read
	//var listNeedToDo ListNeedToDo
	//db.First(&listNeedToDo, 1)
	//db.First(&listNeedToDo, "Name = ?", "Name001")
	//fmt.Println(listNeedToDo, 1)
	//fmt.Println(listNeedToDo, "Name = ?", "Name001")
	//
	//var retrievedTask ListNeedToDo
	//db.First(&retrievedTask, 1)

	// // Update - update product's price to 200
	// db.Model(&listNeedToDo).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&listNeedToDo).Updates(List_Need_To_Do{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&listNeedToDo).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&listNeedToDo, 1)

}
