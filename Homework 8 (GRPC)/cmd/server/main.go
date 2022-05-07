package main

import (
	"grpcProgram/pkg/api"
	"grpcProgram/pkg/reverse"
	"log"
	"net"
	"proto/pkg/reverse"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := reverse.GRPCServer()
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}