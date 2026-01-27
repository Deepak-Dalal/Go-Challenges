/*
The Modern Concurrent "Context-Aware" Port Scanner 🕵️‍♂️
📋 The Requirements:
The Dialer: Instead of the top-level net.Dial, create a net.Dialer struct.

The Context: For every port dial, create a context.WithTimeout. ⏱️

Use a timeout of 500ms.

Use dialer.DialContext(ctx, "tcp", address) to perform the scan.

The Worker Pool: 👷‍♂️

Create a ports channel and a results channel.

Launch 100 workers to process ports 1 to 1024.

The Collection: 📥

The workers should only send open port numbers to the results channel.

Use a separate goroutine to collect these results into a slice so you don't block the workers.
*/
package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	ports:=make(chan int)
	results:=make(chan int)
	var wg sync.WaitGroup
	dialer := net.Dialer{}
	

	for i:=1;i<=100;i++ {
		wg.Add(1)
		go (func (){
			defer wg.Done()
			for port := range ports {
				ctx, cancel:=context.WithTimeout(context.Background(),500*time.Millisecond)
				connection, err:= dialer.DialContext(ctx, "tcp", fmt.Sprintf("127.0.0.1:%d",port))
				cancel()
				if err==nil {
					results<-port
					connection.Close()
				}
			}
		})()
	}
	go func () {
		for i:=1;i<=1024;i++ {
			ports<-i
		}
		close(ports)
	}()
	go  func () {
		wg.Wait()
		close(results)
	}()
	for openPort:=range results {
		fmt.Println(openPort)
	}
}