package main

import (
	"Grpc/Upload/client"
	"Grpc/Upload/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8887", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect server failed", err)
	}
	uploadClient := pb.NewBaseUploadClient(conn)
	clientServer := client.NewClient(uploadClient)
	clientServer.Upload("/usr/local/var/www/Go/src/Grpc/Upload/Go/stroge/ABC.jpeg", "/usr/local/var/www/Go/src/Grpc/Upload/Go/stroge/")
}
