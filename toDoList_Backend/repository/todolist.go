package repository

type ToDoList struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Create_at   string `db:"create_at"`
	Update_at   string `db:"update_at"`
	Delete_at   string `db:"delete_at"`
	Deadline_at string `db:"deadline_at"`
	Done_at string `db:"done_at"`
	Description string `db:"description"`
	Done_in_time int `db:"done_in_time"`
}

type ToDoListRepository interface {
	GetAll() ([]ToDoList, error)
}