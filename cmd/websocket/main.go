package main

import (
	"fmt"
	"log"
	"net/http"
	"treealgos/internal/shared"

	"github.com/gorilla/websocket"
)

// Upgrader config for accepting WebSocket connections
var upgrader = websocket.Upgrader{
	// This function is used to check the origin of the connection.
	// For local development or simplified usage, you can ignore the origin check:
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handleWebSocket is the HTTP handler that upgrades the connection to a WebSocket.
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	shared.Green("Client connected.\n")

	// Continuously read messages from the client
	for {
		// Read message
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			shared.Red("Error reading message: %v\n", err)
			break
		}

		// Log received message
		shared.Faint("Received from client: %s\n", message)

		// Echo the message back to the client with a prefix
		response := fmt.Sprintf("Received: %s\n", string(message))
		if err = conn.WriteMessage(messageType, []byte(response)); err != nil {
			shared.Red("Error writing message: %v\n", err)
			break
		}
	}

	shared.Sred("Client disconnected.\n")
}

func main() {
	// Simple file server to serve HTML/JS client (for testing)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Route for WebSocket connections
	http.HandleFunc("/ws", handleWebSocket)

	// Start HTTP server
	shared.Green("Server started on :8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(shared.Sred("ListenAndServe: ", err))
	}
}
