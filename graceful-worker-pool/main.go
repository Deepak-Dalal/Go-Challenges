package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id       int
	Duration time.Duration //alias int64
}

func worker(jobs <-chan Job,workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job:=range jobs{
		fmt.Printf("Worker %d starting Job %d\n",workerId,job.Id)
		time.Sleep(job.Duration);
		fmt.Printf("Worker %d finished Job %d\n",workerId,job.Id)
	}

}
func main() {
	jobs:=make(chan Job, 10)
	var wg sync.WaitGroup
	wg.Add(3)
	go worker(jobs,1,&wg);
	go worker(jobs,2,&wg);
	go worker(jobs,3,&wg);
	for i:=0;i<10;i++ {
		jobs<-Job{i+1, time.Second}
	}
	close(jobs)
	wg.Wait()
}