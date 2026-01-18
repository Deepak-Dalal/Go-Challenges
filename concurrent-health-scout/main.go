/*
🛠️ Challenge 1: The Concurrent "Health Scout" 🚦
One of Go’s superpowers is how easily it handles multiple tasks at once. Your task is to build a CLI tool that checks if a list of websites is up or down.

📋 The Requirements:
Input: Create a slice of strings containing at least 5–7 different URLs (e.g., https://google.com, https://github.com, etc.). 🌐

Concurrency: Instead of checking them one by one, spawn a goroutine for each URL check. ⚡

Communication: Use a channel to send the results (the URL and its status) back to the main function. 📡

Reporting: Once all checks are done, print a summary to the console.

The "Pro" Touch: Implement a timeout for each request using context or the http.Client timeout settings so one slow website doesn't hang your whole program. ⏱️
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type urlStatus struct {
	url    string
	status string
}

var httpclient = &http.Client{
	Timeout: 5 * time.Second,
}

func checkHealth(url string, ch chan urlStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := httpclient.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	status := "not-ok"
	if err == nil && resp.StatusCode == http.StatusOK {
		status = "ok"
	}
	ch <- urlStatus{url: url, status: status}
}
func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"https://www.baidu.com",
		"https://www.ask.com",
		"https://www.aol.com",
	}
	var ch = make(chan urlStatus, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkHealth(url, ch, &wg)
	}
	wg.Wait()
	for i := 0; i < len(urls); i++ {
		result := <-ch
		fmt.Println(result.url, result.status)
	}
}
