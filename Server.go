package main

import (
	pb "100gombs/protos"
	"io"
	"os"
)

type FileServer struct {
	pb.UnimplementedFileServiceServer
}

const chunkSize = 4096

// const chunkSize = 2048
// const chunkSize = 1024
//const chunkSize = 512
//const chunkSize = 256
//const chunkSize = 128
//const chunkSize = 64
//const chunkSize = 32
//const chunkSize = 16
//const chunkSize = 8

func (f *FileServer) DownloadFile(req *pb.DownloadRequest, stream pb.FileService_DownloadFileServer) error {
	filePath := req.GetPath()
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	buffer := make([]byte, chunkSize)
	var bytesRead int

	for {
		bytesRead, err = file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err = stream.Send(&pb.FileChunk{Data: buffer[:bytesRead]}); err != nil {
			return err
		}
	}

	return nil
}
