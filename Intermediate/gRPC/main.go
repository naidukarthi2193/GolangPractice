package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
