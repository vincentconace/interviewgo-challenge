package health

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

func (h Health) HealthHandler(w http.ResponseWriter, _ *http.Request) {
	h.Status = "OK"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h)
}
