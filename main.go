package main

import (
	"100gombs/client"
	"100gombs/logger"
	pb "100gombs/protos"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	fs := &FileServer{}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, fs)

	logger.InfoLogger.Println("Starting gRPC Server on port 50051")

	go func() {
		if err = s.Serve(lis); err != nil {
			logger.ErrorLogger.Fatalln("Error while running gRPC server: ", err)
		}
	}()

	go func() {
		client.RunClient()
	}()

	select {}
}
