package service

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/entities"
)

type ToDoListService interface {
	GetToDoLists() ([]entities.Todo, error)
}
