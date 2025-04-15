package main

import (
	"flag"
	"log"
	"treealgos/internal/shared"
	"treealgos/internal/stream"
)

func main() {
	addr := flag.String("listenaddr", ":8080", "address where server is listening")
	flag.Parse()

	server := stream.NewServer(*addr)
	shared.Green("Server running on port: http://localhost%v\n", *addr)
	log.Fatal(shared.Sred("%v", server.Start()))
}
