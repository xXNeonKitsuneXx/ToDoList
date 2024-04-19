package service

import (
	"log"

	"github.com/xXNeonKitsuneXx/toDoList_Backend/repository"
)

type toDoListService struct {
	toDoListRepo repository.ToDoListRepository
}

func NewToDoListService(toDoListRepo repository.ToDoListRepository) toDoListService {
	return toDoListService{toDoListRepo: toDoListRepo}
}

func (s toDoListService) GetToDoLists() ([]ToDoListResponse, error) {
	todolists, err := s.toDoListRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	toDoListResponses := []ToDoListResponse{}
	for _, todolist := range todolists {
		toDoListResponse := ToDoListResponse{
			Id: todolist.Id,
			Name: todolist.Name,
			Create_at: todolist.Create_at,
			Update_at: todolist.Update_at,
			Delete_at: todolist.Deadline_at,
			Deadline_at: todolist.Deadline_at,
			Done_at: todolist.Done_at,
			Description: todolist.Description,
			Done_in_time: todolist.Done_in_time,
		}
		toDoListResponses = append(toDoListResponses, toDoListResponse)
	}
	return toDoListResponses, nil
}