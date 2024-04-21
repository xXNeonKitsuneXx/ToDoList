package service

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/entities"
)

type ToDoListService interface {
	GetToDoLists() ([]entities.Todo, error)
}
