package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Wait for server to start
	fmt.Println("Waiting for Chuks server on :9876...")
	for i := 0; i < 50; i++ {
		resp, err := http.Get("http://localhost:9876/")
		if err == nil {
			resp.Body.Close()
			fmt.Println("Server is up!")
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	concurrencies := []int{1, 10, 50, 100}
	duration := 5 * time.Second

	for _, c := range concurrencies {
		runBenchmark("/", "text", c, duration)
	}
	for _, c := range concurrencies {
		runBenchmark("/json", "json", c, duration)
	}
}

func runBenchmark(path, label string, concurrency int, dur time.Duration) {
	var totalRequests int64
	var totalErrors int64
	var totalLatencyNs int64

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: concurrency,
			MaxConnsPerHost:     concurrency,
		},
	}

	var wg sync.WaitGroup
	stop := make(chan struct{})

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			url := fmt.Sprintf("http://localhost:9876%s", path)
			for {
				select {
				case <-stop:
					return
				default:
				}
				start := time.Now()
				resp, err := client.Get(url)
				elapsed := time.Since(start)
				if err != nil {
					atomic.AddInt64(&totalErrors, 1)
					continue
				}
				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()
				if resp.StatusCode == 200 {
					atomic.AddInt64(&totalRequests, 1)
					atomic.AddInt64(&totalLatencyNs, int64(elapsed))
				} else {
					atomic.AddInt64(&totalErrors, 1)
				}
			}
		}()
	}

	time.Sleep(dur)
	close(stop)
	wg.Wait()

	reqs := atomic.LoadInt64(&totalRequests)
	errs := atomic.LoadInt64(&totalErrors)
	latNs := atomic.LoadInt64(&totalLatencyNs)

	rps := float64(reqs) / dur.Seconds()
	var avgLatMs float64
	if reqs > 0 {
		avgLatMs = float64(latNs) / float64(reqs) / 1e6
	}

	fmt.Printf("[%s c=%d] %d reqs in %v | %.0f req/s | avg %.2f ms | %d errors\n",
		label, concurrency, reqs, dur, rps, avgLatMs, errs)
}
