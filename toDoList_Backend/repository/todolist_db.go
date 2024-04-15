package repository

import (
	// "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// So it should be big T or small t ??????
type toDoListRepositoryDB struct {
	db *gorm.DB
}

func NewToDoListRepositoryDB(db *gorm.DB) toDoListRepositoryDB {
	// // Did this is correct way ?
	// db = db.Table("toDoListTable") // Set the table name
	return toDoListRepositoryDB{db: db}
}

func (a toDoListRepositoryDB) GetAll() ([]ToDoList, error) {
	todolists := []ToDoList{}
	result := a.db.Find(&todolists)
	if result.Error != nil {
		return nil, result.Error
	}
	return todolists, nil
}
