package main

import (
	pb "github.com/rendyuwu/go-grpc-server-example/model"
	handler "github.com/rendyuwu/go-grpc-server-example/user/delivery/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalf("error when listening %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, handler.NewUserService())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("error when server %v", err)
	}
}
