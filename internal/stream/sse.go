package stream

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"treealgos/internal/shared"
	"treealgos/internal/treetraversal"
	"treealgos/types/trees"
)

type Server struct {
	listenAddr string
	dataChan   chan trees.MultiChildTreeNode
}

func NewServer(listenaddr string) *Server {
	return &Server{
		listenAddr: listenaddr,
		dataChan:   make(chan trees.MultiChildTreeNode, 10),
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static/stream"))
	se := sendEvents(s.dataChan)
	tree := treeInput(s.dataChan)

	mux.Handle("/", fs)
	mux.Handle("/tree", tree)
	mux.Handle("/events", se)

	return http.ListenAndServe(s.listenAddr, mux)
}

func sendEvents(c chan trees.MultiChildTreeNode) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare HTTP headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		// Flush the headers
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		flusher.Flush()

		// Keep pushing data as it arrives on the channel
		for {
			select {
			case result, ok := <-c:
				if !ok {
					// The channel has been closed. We can log and break the loop,
					// which will end this handler (and close the SSE connection).
					shared.Yellow("Data Channel closed.")
					return
				}
				// Use SSE format => "data: <content>\n\n"
				content := fmt.Sprintf("data: %v\n\n", result.String())
				shared.Faint(content)
				_, _ = w.Write([]byte(content))
				flusher.Flush()

			case <-time.After(10 * time.Second):
				// No data for 10s? Just send a comment or ping to keep the connection alive
				_, _ = w.Write([]byte(": ping\n\n"))
				flusher.Flush()
			}
		}
	})
}

func treeInput(c chan trees.MultiChildTreeNode) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body map[string][]int
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		inputSlice, ok := body["data"]
		if !ok {
			http.Error(w, "Missing 'data' field", http.StatusBadRequest)
			return
		}

		shared.Green("Received post request: %v\n", inputSlice)

		// Build an example root node
		root := &trees.MultiChildTreeNode{
			Val:       1,
			Children:  []*trees.MultiChildTreeNode{},
			IsVisited: false,
			Metadata: trees.TreeMetadata{
				Label: "root",
				Color: shared.Colors[0],
				Depth: 0,
			},
		}

		// Example usage of your TreeBuilder
		value := 2
		counter := &value
		treetraversal.TreeBuilder(root, inputSlice, counter, 1)
		shared.Faint("This should have the built tree already: %v\n", root.String())
		c <- *root

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Tree built and sent to SSE."))
	})
}

func ContentHandler(content string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write([]byte(content))
	}
}
