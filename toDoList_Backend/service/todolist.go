package service

type ToDoListResponse struct {
	Id          int    `json:"column:id"`
	Name        string `json:"column:name"`
	Create_at   string `json:"column:create_at"`
	Update_at   string `json:"column:update_at"`
	Delete_at   string `json:"column:delete_at"`
	Deadline_at string `json:"column:deadline_at"`
	Done_at     string `json:"column:done_at"`
	Description string `json:"column:description"`
	Done_in_time int   `json:"column:done_in_time"`
}

type ToDoListService interface {
	GetToDoLists() ([]ToDoListResponse, error)
}


