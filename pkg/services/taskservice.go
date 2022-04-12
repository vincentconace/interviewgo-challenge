package services

import (
	"fmt"

	"github.com/interviewgo/pkg/data/models"
)

type TaskService struct{}

func New() models.ITaskRepo {
	return &TaskService{}
}

func (ts *TaskService) GetAll() []models.Task {
	return T
}

func (ts *TaskService) Create(task models.Task) (models.Task, error) {
	task.Id = len(T) + 1
	T = append(T, task)
	return task, nil
}

func (ts *TaskService) FindByID(id int) (models.Task, error) {
	for _, task := range T {
		if task.Id == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task with id %d not found", id)
}

func (ts *TaskService) Update(id int, task models.Task) (models.Task, error) {
	for i, task := range T {
		if task.Id == id {
			T[i] = task
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("Task with id %d not found", id)
}

func (ts *TaskService) Delete(id int) error {
	for i, task := range T {
		if task.Id == id {
			T = append(T[:i], T[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task with id %d not found", id)
}
