package mocks

import (
	"github.com/interviewgo/pkg/data/models"
	"github.com/stretchr/testify/mock"
)

type TaskMock struct {
	mock.Mock
}

func (t *TaskMock) GetAll() []models.Task {
	args := t.Called()
	return args.Get(0).([]models.Task)
}

func (t *TaskMock) FindByID(id int) (models.Task, error) {
	args := t.Called(id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (t *TaskMock) Create(task models.Task) (models.Task, error) {
	args := t.Called(task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (t *TaskMock) Update(id int, task models.Task) (models.Task, error) {
	args := t.Called(id, task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (t *TaskMock) Delete(id int) error {
	args := t.Called(id)
	return args.Error(0)
}
