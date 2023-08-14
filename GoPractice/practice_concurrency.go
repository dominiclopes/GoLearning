package main

import (
	"fmt"
	"sync"
	"time"
)

var eJobs = make(chan int)
var eResults = make(chan int)

func eWorker(wg *sync.WaitGroup, i int) {
	for job := range eJobs {
		fmt.Println(i, job)
		eResults <- job
	}
	wg.Done()
}

func eWorkerPool() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go eWorker(&wg, i)
	}
	wg.Wait()
	close(eResults)
}

func eAllocate() {
	for i := 1; i <= 100; i++ {
		eJobs <- i
	}
	close(eJobs)
}

func sum(done chan<- bool) {
	sum := 0
	for result := range eResults {
		sum += result
	}
	fmt.Println("sum is", sum)
	done <- true
}

func main() {
	startTime := time.Now()
	go eAllocate()
	done := make(chan bool)
	go sum(done)
	go eWorkerPool()
	<-done
	endTime := time.Now()
	fmt.Println("time taken", endTime.Sub(startTime).Seconds())

	startTime = time.Now()
	n := 0
	for i := 1; i <= 100; i++ {
		n += i
	}
	fmt.Println(n)
	endTime = time.Now()
	fmt.Println("time taken", endTime.Sub(startTime).Seconds())

}
