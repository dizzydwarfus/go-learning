package stream

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

func SendEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	for {
		content := fmt.Sprintf("Your lucky number: %d\n", rand.IntN(100))

		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// intentional sleep
		time.Sleep(time.Millisecond * 500)
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
