package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Handler handles GET /test-timeout?duration=N
// SSE stream, sends a tick every second
func Handler(w http.ResponseWriter, r *http.Request) {
	duration := 10
	if d, err := strconv.Atoi(r.URL.Query().Get("duration")); err == nil && d >= 1 && d <= 120 {
		duration = d
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, canFlush := w.(http.Flusher)

	start := time.Now()

	// start event
	fmt.Fprintf(w, "event: start\ndata: {\"duration\":%d,\"startTime\":\"%s\"}\n\n", duration, start.Format(time.RFC3339))
	if canFlush {
		flusher.Flush()
	}

	for i := 1; i <= duration; i++ {
		time.Sleep(1 * time.Second)
		elapsed := time.Since(start).Seconds()
		fmt.Fprintf(w, "event: tick\ndata: {\"second\":%d,\"elapsed\":\"%.1fs\",\"remaining\":%d}\n\n", i, elapsed, duration-i)
		if canFlush {
			flusher.Flush()
		}
	}

	total := time.Since(start).Seconds()
	fmt.Fprintf(w, "event: done\ndata: {\"totalElapsed\":\"%.1fs\",\"success\":true}\n\n", total)
	if canFlush {
		flusher.Flush()
	}
}
