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
	// "log"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type List_Need_To_Do struct {
	gorm.Model
	LNTD_ID	   uint `gorm:"primaryKey"`
	Name       string
	Date       time.Time
	Description string
}

func main() {
	dsn := "postgresql://BocchiKitsuNei:DpwPBBlJw78XBb91Au_fpw@bocchikitsunei-8346.8nk.gcp-asia-southeast1.cockroachlabs.cloud:26257/ToDoList_GoLang?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)
	fmt.Println(now)

	// Migrate the schema
	db.AutoMigrate(&List_Need_To_Do{})

	// Create
	db.Create(&List_Need_To_Do{Name: "Name001", Date: now, Description: "Desciption001"})

	// Read
	var listNeedToDo List_Need_To_Do
	db.First(&listNeedToDo, 1)
	db.First(&listNeedToDo, "Name = ?", "Name001")
	fmt.Println(listNeedToDo, 1)
	fmt.Println(listNeedToDo, "Name = ?", "Name001")

	// // Update - update product's price to 200
	// db.Model(&listNeedToDo).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&listNeedToDo).Updates(List_Need_To_Do{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&listNeedToDo).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&listNeedToDo, 1)

	
}
