package main

import (
	"Grpc/Upload/pb"
	"Grpc/Upload/server"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := server.NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterBaseUploadServer(grpcServer, server)
	listener, err := net.Listen("tcp", "127.0.0.1:8887")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}
