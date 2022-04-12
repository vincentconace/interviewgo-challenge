package main

import (
	"log"
	"net/http"

	"github.com/interviewgo/pkg/handlers/health"
	"github.com/interviewgo/pkg/handlers/task"
	"github.com/interviewgo/pkg/services"
)

type Server struct {
	Port string
}

var port = ":8000"

// Create Server
func New() *Server {

	return &Server{
		Port: port,
	}
}

// Run Server
func (s *Server) Run() error {
	log.Println("Sever running on", s.Port)
	routes := s.registerRoutes()
	return http.ListenAndServe(s.Port, routes)
}

// Put here your endpoints
func (s *Server) registerRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	healthController := new(health.Health)
	// End Health

	// Tasks
	taskRepo := services.New()
	taskController := task.New(taskRepo)
	// End Tasks

	// Health Handlers
	mux.HandleFunc("/health", healthController.HealthHandler)
	// End Handlers

	// Tasks
	mux.HandleFunc("/todo/tasks", taskController.GetTasks)
	mux.HandleFunc("/todo/tasks/create", taskController.CreateTask)
	mux.HandleFunc("/todo/tasks/", taskController.FindByID)
	mux.HandleFunc("/todo/tasks/update/", taskController.UpdateTask)
	mux.HandleFunc("/todo/tasks/delete/", taskController.DeleteTask)
	// end Task Controllers

	return mux
}

// This is the point of entry
func main() {
	srv := New()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
