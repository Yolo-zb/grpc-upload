package client

import (
	"Grpc/Upload/pb"
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	service pb.BaseUploadClient
}

func NewClient(service pb.BaseUploadClient) *Client {
	return &Client{
		service: service,
	}
}

func (client *Client) Upload(imagePath string, savePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := client.service.Upload(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}
	fileType := filepath.Ext(imagePath)
	req := &pb.UploadRequest{
		Data: &pb.UploadRequest_FileInfo{
			FileInfo: &pb.FileInfo{
				FileExt:  fileType,
				FilePath: savePath,
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 20000)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}
		req := &pb.UploadRequest{
			Data: &pb.UploadRequest_File{
				File: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	log.Printf("image path : %s", res.GetSavePath())
}
