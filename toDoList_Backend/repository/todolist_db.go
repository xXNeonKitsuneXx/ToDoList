package repository

import (
	// "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type toDoListRepositoryDB struct {
	db *gorm.DB
}

func NewToDoListRepositoryDB(db *gorm.DB) toDoListRepositoryDB {
	return toDoListRepositoryDB{db: db}
}

func (r toDoListRepositoryDB) GetAll() ([]ToDoList, error) {
	todolists := []ToDoList{}
	result := r.db.Find(&todolists)
	if result.Error != nil {
		return nil, result.Error
	}
	return todolists, nil
}
