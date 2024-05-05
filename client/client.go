package client

import (
	"100gombs/logger"
	pb "100gombs/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

var (
	client pb.FileServiceClient
)

func RunClient() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		logger.ErrorLogger.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a FileService client
	client = pb.NewFileServiceClient(conn)

	if err = run(); err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}
