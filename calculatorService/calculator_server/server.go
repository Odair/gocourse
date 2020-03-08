package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

)

type server struct{}

func main() {
	fmt.Println("Hello Calculator")

	listen, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
