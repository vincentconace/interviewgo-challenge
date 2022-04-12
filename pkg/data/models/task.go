package models

// Task entity
type Task struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Repositori interface
type ITaskRepo interface {
	GetAll() []Task
	FindByID(id int) (Task, error)
	Create(task Task) (Task, error)
	Update(id int, task Task) (Task, error)
	Delete(id int) error
}
