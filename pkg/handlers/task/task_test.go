package task_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/interviewgo/pkg/data/models"
	"github.com/interviewgo/pkg/handlers/task"
	"github.com/interviewgo/pkg/mocks"
	"gotest.tools/assert"
)

// fixtures
var (
	simpleTask = models.Task{
		Id:     3,
		Name:   "A simple task",
		Status: "In progress",
	}
	postTask = models.Task{
		Name:   "A simple task",
		Status: "In progress",
	}
	updateTask = models.Task{
		Status: "Done",
	}
	emptyTask      = models.Task{}
	emptyTaskList  = []models.Task{}
	incompleteTask = models.Task{
		Name: "A simple task",
	}
	taskwithoutstatus = models.Task{
		Name:   "A simple task",
		Status: " ",
	}
	updatedTask = models.Task{
		Id:     3,
		Name:   "A simple task",
		Status: "Done",
	}
	okTasks = []models.Task{
		{
			Id:     1,
			Name:   "Análisis de Requerimiento",
			Status: "In progress",
		},
		{
			Id:     2,
			Name:   "Creación de Endpoints",
			Status: "Waiting for Response",
		},
	}
)

func TestGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/todo/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Happy Path", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		mockTaskRepo.On("GetAll").Return(okTasks)

		taskHandler := task.New(&mockTaskRepo)
		rr := httptest.NewRecorder()
		tasksController := http.HandlerFunc(taskHandler.GetTasks)
		tasksController.ServeHTTP(rr, req)

		dataResponse := rr.Body.Bytes()
		response := make([]models.Task, 0)
		json.Unmarshal(dataResponse, &response)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, response[0].Id, okTasks[0].Id)
		assert.Equal(t, response[0].Name, okTasks[0].Name)
		assert.Equal(t, response[0].Status, okTasks[0].Status)
	})

	t.Run("When returns empty data", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		mockTaskRepo.On("GetAll").Return(emptyTaskList)

		taskHandler := task.New(&mockTaskRepo)
		rr := httptest.NewRecorder()
		tasksController := http.HandlerFunc(taskHandler.GetTasks)
		tasksController.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		// read data
		dataResponse := rr.Body.Bytes()
		response := make([]models.Task, 0)
		json.Unmarshal(dataResponse, &response)
		assert.Equal(t, len(response), 0)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		mockTaskRepo.On("Create", postTask).Return(simpleTask, nil)
		tController := task.New(&mockTaskRepo)
		body, _ := json.Marshal(postTask)
		req, err := http.NewRequest("POST", "/todo/tasks/create", bytes.NewReader(body))
		req.Header.Set("content-type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.CreateTask)
		responseTask := new(models.Task)
		taskController.ServeHTTP(rr, req)
		err = json.NewDecoder(rr.Body).Decode(&responseTask)
		defer req.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		//En mi opinion verificaria antes el codigo de respuesta que
		//el error de Unmarshalling
		//assert.Equal(t, rr.Code, http.StatusCreated)
		assert.NilError(t, err)
		assert.Equal(t, rr.Code, http.StatusCreated)
		assert.Equal(t, responseTask.Id, simpleTask.Id)
		assert.Equal(t, responseTask.Name, simpleTask.Name)
		assert.Equal(t, responseTask.Status, simpleTask.Status)
	})

	t.Run("When you sent incomplete data", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		mockTaskRepo.On("Create", incompleteTask).Return(emptyTask, errors.New("name or status not valid"))
		tController := task.New(&mockTaskRepo)
		body, _ := json.Marshal(incompleteTask)
		req, err := http.NewRequest("POST", "/todo/tasks/create", bytes.NewReader(body))
		req.Header.Set("content-type", "application/json")
		defer req.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.CreateTask)
		taskController.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})

	t.Run("When you sent empty data", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		mockTaskRepo.On("Create", emptyTask).Return(emptyTask, errors.New("name or status not valid"))
		tController := task.New(&mockTaskRepo)
		body, _ := json.Marshal(emptyTask)
		req, err := http.NewRequest("POST", "/todo/tasks/create", bytes.NewReader(body))
		req.Header.Set("content-type", "application/json")
		defer req.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.CreateTask)
		taskController.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		fixtureID := 3
		mockTaskRepo.On("Update", fixtureID, updateTask).Return(updatedTask, nil)
		tController := task.New(&mockTaskRepo)
		body, _ := json.Marshal(updateTask)
		req, err := http.NewRequest("PUT", "/todo/tasks/update/3", bytes.NewReader(body))
		req.Header.Set("content-type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.UpdateTask)
		taskController.ServeHTTP(rr, req)
		responseTask := new(models.Task)
		err = json.NewDecoder(rr.Body).Decode(&responseTask)
		assert.NilError(t, err)
		assert.Equal(t, rr.Code, http.StatusAccepted)
		assert.Equal(t, responseTask.Id, updatedTask.Id)
		assert.Equal(t, responseTask.Name, updatedTask.Name)
		assert.Equal(t, responseTask.Status, updatedTask.Status)
	})

	t.Run("When id doesn't exist", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		fixtureID := 5
		mockTaskRepo.On("Update", fixtureID, updateTask).Return(emptyTask, errors.New("id param is incorrect"))
		tController := task.New(&mockTaskRepo)
		body, _ := json.Marshal(updateTask)
		req, err := http.NewRequest("PUT", "/todo/tasks/update/5", bytes.NewReader(body))
		req.Header.Set("content-type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.UpdateTask)
		taskController.ServeHTTP(rr, req)
		defer req.Body.Close()
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		fixtureID := 1
		mockTaskRepo.On("Delete", fixtureID).Return(nil)
		tController := task.New(&mockTaskRepo)
		req, err := http.NewRequest("DELETE", "/todo/tasks/delete/1", nil)
		req.Header.Set("content-type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.DeleteTask)
		taskController.ServeHTTP(rr, req)
		assert.NilError(t, err)
		assert.Equal(t, rr.Code, http.StatusAccepted)
	})

	t.Run("When id doesn't exist", func(t *testing.T) {
		var mockTaskRepo mocks.TaskMock
		fixtureID := 8
		mockTaskRepo.On("Delete", fixtureID).Return(errors.New("invalid id or repo is empty"))
		tController := task.New(&mockTaskRepo)
		req, err := http.NewRequest("DELETE", "/todo/tasks/delete/8", nil)
		req.Header.Set("content-type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		taskController := http.HandlerFunc(tController.DeleteTask)
		taskController.ServeHTTP(rr, req)
		assert.NilError(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
}
