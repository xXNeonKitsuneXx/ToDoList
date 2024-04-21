package repository

import (
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/entities"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/internal/utils/v"
	"time"
)

type toDoListRepositoryMock struct {
	todolists []entities.Todo
}

func NewToDoListRepositoryMock() toDoListRepositoryMock {
	todolists := []entities.Todo{
		{ID: v.Ptr(uint(111)), Name: v.Ptr("R"),
			CreatedAt:   v.Ptr(time.Date(2024, 4, 19, 22, 20, 6, 0, time.UTC)),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			DeadlineAt:  v.Ptr(time.Date(2024, 4, 20, 22, 20, 7, 0, time.UTC)),
			DoneAt:      v.Ptr(time.Date(2024, 4, 20, 22, 10, 7, 0, time.UTC)),
			Description: v.Ptr("AAA111"),
			DoneInTime:  new(bool),
		},
		{ID: v.Ptr(uint(222)), Name: v.Ptr("E"),
			CreatedAt:   v.Ptr(time.Date(2024, 4, 19, 22, 20, 6, 0, time.UTC)),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			DeadlineAt:  v.Ptr(time.Date(2024, 4, 20, 22, 20, 7, 0, time.UTC)),
			DoneAt:      v.Ptr(time.Date(2024, 4, 20, 22, 10, 7, 0, time.UTC)),
			Description: v.Ptr("222BBB"),
			DoneInTime:  new(bool),
		},
		{ID: v.Ptr(uint(333)), Name: v.Ptr("N"),
			CreatedAt:   v.Ptr(time.Date(2024, 4, 19, 22, 20, 6, 0, time.UTC)),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			DeadlineAt:  v.Ptr(time.Date(2024, 4, 20, 22, 20, 7, 0, time.UTC)),
			DoneAt:      nil,
			Description: nil,
			DoneInTime:  nil,
		},
	}
	return toDoListRepositoryMock{todolists: todolists}
}

func (r toDoListRepositoryMock) GetAll() ([]entities.Todo, error) {
	return r.todolists, nil
}
