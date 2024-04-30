package dtos

import "time"

type ToDoListResponse struct {
	ID          *uint      `json:"id" validate:"required"`
	Name        *string    `json:"name" validate:"required"`
	CreatedAt   *time.Time `json:"created_at" validate:"required"`
	UpdatedAt   *time.Time `json:"updated_at" validate:"required"`
	DeadlineAt  *time.Time `json:"deadline_at" validate:"required"`
	DoneAt      *time.Time `json:"done_at" validate:"required"`
	Description *string    `json:"description" validate:"required"`
	DoneInTime  *bool      `json:"done_in_time" validate:"required"`
}
