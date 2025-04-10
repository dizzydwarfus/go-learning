package main

import (
	"flag"
	"log"
	"net/http"
	"treealgos/internal/shared"
	"treealgos/internal/stream"
)

type Server struct {
	listenAddr string
}

func main() {
	addr := flag.String("listenaddr", ":8000", "address where server is listening")
	flag.Parse()

	server := &Server{listenAddr: *addr}
	shared.Green("Server running on port: http://localhost%v\n", *addr)
	log.Fatal(shared.Sred("%v", server.Start()))
}

func (s *Server) Start() error {
	http.HandleFunc("/", stream.Home)
	http.HandleFunc("/events", stream.SendEvents)
	return http.ListenAndServe(s.listenAddr, nil)
}
