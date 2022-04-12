package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/interviewgo/pkg/data/models"
)

type TaskRequest struct {
	taskservices models.ITaskRepo
}

type ITaskRequest interface {
	GetTasks(w http.ResponseWriter, _ *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

func New(repo models.ITaskRepo) ITaskRequest {
	return &TaskRequest{
		taskservices: repo,
	}
}

// GetTasks returns all tasks
func (tr *TaskRequest) GetTasks(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(tr.taskservices.GetAll())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// CreateTask creates a new task
func (tr *TaskRequest) CreateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var t models.Task
	err = json.Unmarshal(bodyBytes, &t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if t.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if t.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := tr.taskservices.Create(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(data))
}

// FindByID returns a task by id
func (tr *TaskRequest) FindByID(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/todo/tasks/")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, err := tr.taskservices.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}

// UpdateTask updates a task
func (tr *TaskRequest) UpdateTask(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/todo/tasks/update/")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var task models.Task
	err = json.Unmarshal(bodyBytes, &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err = tr.taskservices.Update(id, task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, string(data))
}

// DeleteTask deletes a task
func (tr *TaskRequest) DeleteTask(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/todo/tasks/delete/")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = tr.taskservices.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
