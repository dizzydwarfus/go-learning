package stream

import (
	"fmt"
	"net/http"
	"time"
)

func SendEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"Testing", "how", "data", "streaming", "works", "in", "go"}

	for _, token := range tokens {
		content := fmt.Sprintf("test text: %s\n", string(token))

		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// intentional sleep
		time.Sleep(time.Millisecond * 1000)
	}
}

func ContentHandler(content string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write([]byte(content))
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go to /events to get streaming data")
}
