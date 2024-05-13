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
	ErrorLogger          = log.New(os.Stdout, "[ERROR]: \t", log.Lshortfile|log.Ldate|log.Ltime)
	InfoLogger           = log.New(os.Stdout, "[INFO]: \t", log.Lshortfile|log.Ldate|log.Ltime)
	totalBytes     int64 = 1024 * 1024 * 10
	totalBytesLock sync.Mutex
	stopLogging    = make(chan struct{})
	counter        = 1
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

// download10Gb handler is the main function to download file from server
// it also updates statistics in statistics table:
// total_downloads, megabits_transferred, concurrent_downloads, and average_download_time
func download10Gb(c *gin.Context) {
	start := time.Now().Unix()

	file, err := os.Open("../files/10GB.bin")
	if err != nil {
		logger.ErrorLogger.Fatalf("error while opening file: %v", err)
	}
	var wg sync.WaitGroup
	defer file.Close()

	chunkChan := make(chan []byte)

	//go loggingThroughput()

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

	duration := time.Now().Unix() - start
	logger.InfoLogger.Println(" Download Duration: \t", duration)
	//
	//err = updateFileInfoAfterDownload("10GB.bin", duration)
	//if err != nil {
	//	logger.ErrorLogger.Println(err)
	//	return
	//}

	stopLogging <- struct{}{}
}

func download1Gb(c *gin.Context) {
	start := time.Now().Unix()

	file, err := os.Open("../files/1GB.bin")
	if err != nil {
		logger.ErrorLogger.Fatalf("error while opening file: %v", err)
	}
	var wg sync.WaitGroup
	defer file.Close()

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

	duration := time.Now().Unix() - start
	logger.InfoLogger.Println(" Download Duration: \t", duration)

	err = updateFileInfoAfterDownload("10GB.bin", duration)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	stopLogging <- struct{}{}
}

// getTotalDownloads returns total amount of all downloads
func getTotalDownloads(c *gin.Context) {
	v, err := database.GetTotalDownloads()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"amount": v,
	})
}
