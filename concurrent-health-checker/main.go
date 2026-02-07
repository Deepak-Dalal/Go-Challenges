package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL     string
	Status  int
	Latency time.Duration
	Err     error
}

type Reporter interface {
	Report(r Result)
}

type ConsoleReporter struct {}

func (cr *ConsoleReporter) Report(r Result){
	if r.Err != nil {
        fmt.Printf("❌ ERROR: %s | %v\n", r.URL, r.Err)
        return
    }
    
    statusIcon := "✅"
    if r.Status >= 400 {
        statusIcon = "⚠️"
    }
    
    fmt.Printf("%s %d | %10s | %s\n", statusIcon, r.Status, r.Latency, r.URL)
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

	urlsToCheck:=make(chan string)
	resultsChan:=make(chan Result)
	var wg sync.WaitGroup

	httpClient:=&http.Client{
		Timeout: 10*time.Second,
	}

	for i:=0;i<5;i++ {
		wg.Add(1)
		go func (){
			defer wg.Done()
			for url:= range urlsToCheck {
				result:=Result{}
				ctx,cancel:=context.WithTimeout(context.Background(), 3*time.Second)
				
				
				req, err:=http.NewRequestWithContext(ctx,http.MethodGet,url,nil)
				if err!=nil {
					result.Err=err
					resultsChan<-result
					fmt.Println("Error in creating request")
					cancel()
					continue
				}
				startTime:=time.Now()
				resp, err:=httpClient.Do(req)
				latency:=time.Since(startTime)
				result.Latency=latency
				if err != nil {
					result.Err=err
					resultsChan<-result
					fmt.Println("Error in executing request")
					cancel()
					continue
				}
				if resp!=nil {
					result.Status=resp.StatusCode
					resp.Body.Close()
				}
				
				resultsChan<-result
				cancel()
			}
		}()
	}

	go func() {
		for _,url:=range urls {
			urlsToCheck<-url
		}
		close(urlsToCheck)
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	downCount:=0
	upCount:=0
	var r Reporter = &ConsoleReporter{}
	for result:=range resultsChan {
		if result.Err == nil && result.Status == http.StatusOK{
			upCount++;
		} else {
			downCount++;
		}
		r.Report(result)
	}
	fmt.Printf("%d services up, %d services down",upCount,downCount)
}