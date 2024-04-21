package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/dtos"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/service"
)

type toDoListHandler struct {
	toDoListSer service.ToDoListService
}

func NewToDoListHandler(toDoListSer service.ToDoListService) toDoListHandler {
	return toDoListHandler{toDoListSer: toDoListSer}
}

func (h *toDoListHandler) GetToDoLists(c *fiber.Ctx) error {

	toDoListsResponse := make([]dtos.ToDoListResponse, 0)

	todolists, err := h.toDoListSer.GetToDoLists()

	if err != nil {
		return err
	}

	for _, todolist := range todolists {
		toDoListsResponse = append(toDoListsResponse, dtos.ToDoListResponse{
			ID:          todolist.ID,
			Name:        todolist.Name,
			CreatedAt:   todolist.CreatedAt,
			UpdatedAt:   todolist.UpdatedAt,
			DeletedAt:   todolist.DeadlineAt,
			DeadlineAt:  todolist.DeadlineAt,
			DoneAt:      todolist.DoneAt,
			Description: todolist.Description,
			DoneInTime:  todolist.DoneInTime,
		})
	}
	return c.JSON(toDoListsResponse)
}
