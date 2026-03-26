package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// Handler handles GET /test-timeout
// Sleep 15s to exceed the 10s maxDuration limit
func Handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(15 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "timeout test completed",
		"elapsed": time.Since(start).Milliseconds(),
	})
}
