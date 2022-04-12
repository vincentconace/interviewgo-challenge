package services

import (
	"fmt"

	"github.com/interviewgo/pkg/data/models"
)

type TaskService struct{}

// Create a new ITaskRepo implementation
func New() models.ITaskRepo {
	return &TaskService{}
}

// GetAll returns all tasks
func (ts *TaskService) GetAll() []models.Task {
	return T
}

// Create returns a new task
func (ts *TaskService) Create(task models.Task) (models.Task, error) {
	task.Id = len(T) + 1
	T = append(T, task)
	return task, nil
}

// FindByID returns a task by id
func (ts *TaskService) FindByID(id int) (models.Task, error) {
	for _, task := range T {
		if task.Id == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task with id %d not found", id)
}

// Update returns a task by id
func (ts *TaskService) Update(id int, task models.Task) (models.Task, error) {
	for i, k := range T {
		if k.Id == id {
			task.Id = k.Id
			T[i] = task
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task with id %d not found", id)
}

// Delete returns a task by id
func (ts *TaskService) Delete(id int) error {
	for i, task := range T {
		if task.Id == id {
			T = append(T[:i], T[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task with id %d not found", id)
}
