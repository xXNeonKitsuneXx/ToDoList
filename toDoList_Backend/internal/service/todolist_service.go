package service

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/entities"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/repository"
	"log"
)

type toDoListService struct {
	toDoListRepo repository.ToDoListRepository
}

func NewToDoListService(toDoListRepo repository.ToDoListRepository) toDoListService {
	return toDoListService{toDoListRepo: toDoListRepo}
}

func (s toDoListService) GetToDoLists() ([]entities.Todo, error) {
	todolists, err := s.toDoListRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	toDoListResponses := []entities.Todo{}
	for _, todolist := range todolists {
		toDoListResponse := entities.Todo{
			ID:          todolist.ID,
			Name:        todolist.Name,
			CreatedAt:   todolist.CreatedAt,
			UpdatedAt:   todolist.UpdatedAt,
			DeletedAt:   todolist.DeletedAt,
			DeadlineAt:  todolist.DeadlineAt,
			DoneAt:      todolist.DoneAt,
			Description: todolist.Description,
			DoneInTime:  todolist.DoneInTime,
		}
		toDoListResponses = append(toDoListResponses, toDoListResponse)
	}
	return toDoListResponses, nil
}
