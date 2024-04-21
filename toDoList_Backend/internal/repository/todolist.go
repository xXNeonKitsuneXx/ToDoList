package repository

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/entities"
)

type ToDoListRepository interface {
	GetAll() ([]entities.Todo, error)
}
