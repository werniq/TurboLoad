package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/werniq/turboload/logger"
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

const chunkSize = 1024 * 1024

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
			throughput := float64(totalBytes-prevTotalBytes) / (1024 * 1024)
			logger.InfoLogger.Printf("Throughput: %.2f MB/s\n", throughput)
			prevTotalBytes = totalBytes
			totalBytesLock.Unlock()
		}
	}
}

func download10Gb(c *gin.Context) {
	file, err := os.Open("./files/10GB.bin")
	if err != nil {
		logger.ErrorLogger.Fatalf("error while opening file: %v", err)
	}
	var wg sync.WaitGroup
	defer file.Close()

	// Create a channel to communicate between goroutines
	chunkChan := make(chan []byte)

	go loggingThroughput()

	// Concurrently read from file
	wg.Add(1)
	go func() {
		defer close(chunkChan)
		buffer := make([]byte, chunkSize)

		for {
			bytesRead, err := file.Read(buffer)
			if err == io.EOF {
				break
			}

			if err != nil {
				logger.ErrorLogger.Fatalf("error while reading chunk: %v", err)
			}
			chunkChan <- buffer[:bytesRead]
		}
		wg.Done()
	}()

	// concurrently write to response writer
	c.Header("Content-Disposition", "attachment; filename=10GB.bin")
	c.Header("Content-Type", "application/octet-stream")

	wg.Add(1)
	go func() {
		// write file data directly to the response writer
		defer wg.Done()
		for chunk := range chunkChan {
			_, err = c.Writer.Write(chunk)
			if err != nil {
				logger.ErrorLogger.Fatalf("error while writing chunk to response: %v", err)
			}
			// Update total bytes
			totalBytesLock.Lock()
			totalBytes += int64(len(chunk))
			totalBytesLock.Unlock()
		}
	}()

	wg.Wait()
	stopLogging <- struct{}{}
}
