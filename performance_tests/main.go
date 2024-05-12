package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	RequestsAmount = 100
	//RequestsAmount = 10
	//RequestsAmount = 1
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < RequestsAmount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			resp, err := http.Get("http://localhost:8080/download-10-gb")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			// Create a buffer to reuse for reading the response body
			buffer := make([]byte, 1024)
			var totalBytes int
			for {
				n, err := resp.Body.Read(buffer)
				if err != nil {
					if err != io.EOF {
						panic(err)
					}
					break
				}
				totalBytes += n
			}

			fmt.Printf("Downloaded %d bytes\n", totalBytes)
			time.Sleep(time.Second * 5)
		}()
	}
	wg.Wait()
}
