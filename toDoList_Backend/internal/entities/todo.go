package entities

import "time"

type Todo struct {
	ID          *uint
	Name        *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeadlineAt  *time.Time
	DoneAt      *time.Time
	Description *string
	DoneInTime  *bool
}
