package repository

type ToDoList struct {
	Id          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Create_at   string `gorm:"column:create_at"`
	Update_at   string `gorm:"column:update_at"`
	Delete_at   string `gorm:"column:delete_at"`
	Deadline_at string `gorm:"column:deadline_at"`
	Done_at     string `gorm:"column:done_at"`
	Description string `gorm:"column:description"`
	Done_in_time int    `gorm:"column:done_in_time"`
}

type ToDoListRepository interface {
	GetAll() ([]ToDoList, error)
}