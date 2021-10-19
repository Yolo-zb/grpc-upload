package server

import (
	"Grpc/Upload/pb"
	"bytes"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.BaseUploadServer
	ImageStore ImageStore
}

func NewServer() *Server {
	return &Server{
		ImageStore: ImageStore{},
	}
}

func (server *Server) Upload(stream pb.BaseUpload_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	imageData := bytes.Buffer{} // byte类型的缓冲器，用来存储流消息的文件数据
	// 接收流消息直至接收完全部
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}
		chunk := req.GetFile()
		size := len(chunk)
		log.Printf("received a chunk with size: %d", size)
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}
	// 接收完了，开始写入
	imageID, err := server.ImageStore.Save(imageData, req.GetFileInfo().GetFilePath(), req.GetFileInfo().FileExt)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}
	res := &pb.UploadResponse{
		SavePath: imageID,
	}
	// 返回结果给cilent端并关闭链接
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
