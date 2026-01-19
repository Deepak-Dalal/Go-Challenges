/*
The "Atomic Rate Limiter" 🛡️
In high-traffic systems, you often need to track things across goroutines without the overhead of a full Mutex. For this challenge, we are going to dive into Low-level Synchronization using the sync/atomic package and Signaling.

📋 The Scenario:
You are building a high-speed middleware that counts how many requests are being processed. If the count exceeds a certain limit, it should "reject" the request.

📋 The Requirements:
The Counter: Use an int64 variable to track "Active Requests". 🔢

The Simulator: * Launch 50 goroutines almost at once.

Each goroutine represents an incoming request.

The Logic:

When a request starts, it should increment the counter atomically. ⚛️

If the counter's value (after incrementing) is greater than 10, the request is "Rejected" (print a message and decrement the counter immediately).

If the counter is 10 or less, the request is "Accepted". It should sleep for 100ms to simulate work, then decrement the counter atomically before finishing.

The Report: After all 50 goroutines finish, print the final value of the counter. (It should be 0!) 🏁
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var activeRequestsCounter int64
	var wg sync.WaitGroup

	for i:=0;i<50;i++ {
		wg.Add(1)
		go (func () {
			defer wg.Done()
			currentCount:=atomic.AddInt64(&activeRequestsCounter,1);
			if currentCount>10 {
				fmt.Println("Request rejected by rate limiter")
				atomic.AddInt64(&activeRequestsCounter,-1)
			} else {
				time.Sleep(100*time.Millisecond)
				atomic.AddInt64(&activeRequestsCounter,-1);
			}
		})()
	}
	wg.Wait()

	fmt.Println("Counter value: ", atomic.LoadInt64(&activeRequestsCounter))
}