package repository

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/entities"
)

type ToDoListRepository interface {
	GetAll() ([]entities.Todo, error)
}
