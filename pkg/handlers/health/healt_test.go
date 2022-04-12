package health_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	hh "github.com/interviewgo/pkg/handlers/health"
	"gotest.tools/assert"
)

func TestHealthHandler(t *testing.T) {
	t.Run("Check status server", func(t *testing.T) {
		expected := hh.Health{"OK"}
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		tasksController := http.HandlerFunc(expected.HealthHandler)
		tasksController.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		dataResponse := rr.Body.Bytes()
		response := new(hh.Health)
		json.Unmarshal(dataResponse, &response)
		assert.Equal(t, response.Status, expected.Status)
	})
}

func simulateServer() {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		log.Fatal(err)
	}
	h := hh.Health{}

	rr := httptest.NewRecorder()
	tasksController := http.HandlerFunc(h.HealthHandler)
	tasksController.ServeHTTP(rr, req)
}

func BenchmarkHealthHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulateServer()
	}
}
