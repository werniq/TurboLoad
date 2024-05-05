package client

import (
	"100gombs/logger"
	pb "100gombs/protos"
	"100gombs/utils"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var (
	ErrorLogger    = log.New(os.Stdout, "[ERROR]: \t", log.Lshortfile|log.Ldate|log.Ltime)
	InfoLogger     = log.New(os.Stdout, "[INFO]: \t", log.Lshortfile|log.Ldate|log.Ltime)
	totalBytes     int64
	totalBytesLock sync.Mutex
	stopLogging    = make(chan struct{})
	counter        = 1
	result         = []float64{}
)

func run() error {
	r := gin.Default()
	r.Use(cors.Default())
	applyRoutes(r)
	return r.Run(":8080")
}

func loggingThroughput() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var prevTotalBytes int64

	for {
		select {
		case <-stopLogging:
			return
		case <-ticker.C:
			totalBytesLock.Lock()
			result = append(result, float64(totalBytes-prevTotalBytes)/(1024*1024))
			prevTotalBytes = totalBytes
			totalBytesLock.Unlock()
		}
	}
}

func download10Gb(c *gin.Context) {
	downloadRequest := &pb.DownloadRequest{Path: "./1GB.bin"}

	stream, err := client.DownloadFile(context.Background(), downloadRequest)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to download file: %v\n", err)
	}

	// setting response headers
	c.Header("Content-Disposition", "attachment; filename=10GB.bin")
	c.Header("Content-Type", "application/octet-stream")

	var chunk = &pb.FileChunk{}

	// Write file data directly to the response writer
	go loggingThroughput()

	// Set response headers
	c.Header("Content-Disposition", "attachment; filename=10GB.bin")
	c.Header("Content-Type", "application/octet-stream")

	// Write file data directly to the response writer
	for {
		chunk, err = stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			logger.ErrorLogger.Fatalf("error while receiving chunk: %v", err)
		}

		// Write chunk data to response
		_, err = c.Writer.Write(chunk.Data)
		if err != nil {
			logger.ErrorLogger.Fatalf("error while writing chunk to response: %v", err)
		}

		// Update total bytes
		totalBytesLock.Lock()
		totalBytes += int64(len(chunk.Data))
		totalBytesLock.Unlock()
	}

	logger.InfoLogger.Println("File sent successfully")
	stopLogging <- struct{}{}

	logger.InfoLogger.Println("Average Throughput: ", utils.Avg(result), "MB/s")
}
