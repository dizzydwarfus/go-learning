package stream

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"treealgos/internal/shared"
	"treealgos/internal/treetraversal"
	"treealgos/types/trees"

	"github.com/google/uuid"
)

type Server struct {
	listenAddr string
}

type SSEHub struct {
	mu      sync.Mutex
	streams map[string]chan trees.MultiChildTreeNode
}

func NewSSEHub() *SSEHub {
	return &SSEHub{
		streams: make(map[string]chan trees.MultiChildTreeNode),
	}
}

func (hub *SSEHub) GetSessionChannel(sessionId string) chan trees.MultiChildTreeNode {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	ch, exists := hub.streams[sessionId]
	if !exists {
		ch = make(chan trees.MultiChildTreeNode, 50)
		hub.streams[sessionId] = ch
	}
	return ch
}

func (hub *SSEHub) CloseSessionChannel(sessionId string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	if ch, ok := hub.streams[sessionId]; ok {
		close(ch)
		delete(hub.streams, sessionId)
	}
}

func (hub *SSEHub) Publish(sessionId string, data trees.MultiChildTreeNode) {
	ch := hub.GetSessionChannel(sessionId)
	ch <- data
}

func NewServer(listenaddr string) *Server {
	return &Server{
		listenAddr: listenaddr,
	}
}

func (s *Server) Start() error {
	hub := NewSSEHub()
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static/stream"))
	se := sendEvents(hub)
	tree := treeInput(hub)

	mux.Handle("/", fs)
	mux.HandleFunc("/init", initHandler)
	mux.Handle("/tree", tree)
	mux.Handle("/events", se)

	return http.ListenAndServe(s.listenAddr, mux)
}

func sendEvents(hub *SSEHub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Missing session cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionId := cookie.Value
		// Prepare HTTP headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Flush the headers
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		flusher.Flush()

		ch := hub.GetSessionChannel(sessionId)
		// Keep pushing data as it arrives on the channel
		for {
			select {
			case data, ok := <-ch:
				if !ok {
					// The channel was closed. End this SSE connection.
					log.Printf(shared.Syellow("Channel closed for session %s"), sessionId)
					return
				}
				// SSE format: "data: <JSON or string>\n\n"
				jsonData := treetraversal.ShowJSONTree(&data)
				fmt.Fprintf(w, "data: %s\n\n", string(jsonData))
				flusher.Flush()

			case <-time.After(5 * time.Second):
				// keep-alive ping to avoid timeouts
				fmt.Fprintf(w, ": ping\n\n")
				flusher.Flush()

			case <-r.Context().Done():
				// Client disconnected
				log.Printf(shared.Sred("Client disconnected for session %s"), sessionId)
				// Optionally close & remove channel if no one else will use it
				hub.CloseSessionChannel(sessionId)
				return
			}
		}
	})
}

func treeInput(hub *SSEHub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		cookie, err := r.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Missing session cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessionId := cookie.Value
		var body map[string][]int

		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		inputSlice, ok := body["data"]
		if !ok {
			http.Error(w, "Missing 'data' field", http.StatusBadRequest)
			return
		}

		log.Printf(shared.Sfaint("Received post request: %v\n"), inputSlice)

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
		hub.Publish(sessionId, *root) // sends to the channel for that session

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK. Data published to session %s\n", sessionId)
	})
}

func ContentHandler(content string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write([]byte(content))
	}
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	// If no cookie, set a new one
	_, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		sessionId := uuid.NewString()
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sessionId,
			Path:  "/",
		})
	}
	fmt.Fprintln(w, "OK, cookie set if missing")
}
