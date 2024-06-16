package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	var (
		url                string
		concurrentRequests int
		timeout            time.Duration
	)
	flag.StringVar(&url, "url", "http://localhost:8000/long", "rest uri to test")
	flag.IntVar(&concurrentRequests, "requests", 10, "concurrent request number")
	flag.DurationVar(&timeout, "timeout", 15*time.Second, "timeout for request")
	flag.Parse()

	var wg sync.WaitGroup
	statusCodeCounter := make(map[int]int)
	mu := &sync.Mutex{}

	client := &http.Client{
		Timeout: timeout,
	}

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			response, err := client.Get(url)
			if err != nil {
				fmt.Println("error:", err.Error())
				mu.Lock()
				statusCodeCounter[500]++
				mu.Unlock()
				return
			}
			defer response.Body.Close()

			mu.Lock()
			statusCodeCounter[response.StatusCode]++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("\nHTTP status counter:")
	for code, count := range statusCodeCounter {
		fmt.Printf("Code %d: %d times \n", code, count)
	}
}
