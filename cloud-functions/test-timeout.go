package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Handler handles GET /test-timeout?duration=N
// Sleeps for the specified duration then returns elapsed time
func Handler(w http.ResponseWriter, r *http.Request) {
	duration := 10
	if d, err := strconv.Atoi(r.URL.Query().Get("duration")); err == nil && d >= 1 && d <= 120 {
		duration = d
	}

	start := time.Now()
	time.Sleep(time.Duration(duration) * time.Second)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "timeout test completed",
		"duration": duration,
		"elapsed":  time.Since(start).Milliseconds(),
	})
}
