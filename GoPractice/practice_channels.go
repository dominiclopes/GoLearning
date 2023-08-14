package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ----------------------------------------
// func main() {
// 	var a chan int

// 	if a == nil {
// 		fmt.Printf("type of channel a is %T\n", a)
// 		fmt.Println("channel a is nil")
// 		a = make(chan int)
// 		fmt.Printf("type of channel a is %T\n", a)
// 	}
// }

// ----------------------------------------
// func hello(done chan<- bool) {
// 	fmt.Println("hello world go routine")
// 	done <- true
// }

// func main() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<-done
// 	fmt.Println("main function")

// }

// ----------------------------------------
// func digits(number int, digitChan chan<- int) {
// 	for number != 0 {
// 		digit := number % 10
// 		digitChan <- digit
// 		number /= 10
// 	}
// 	close(digitChan)
// }

// func calcSquares(number int, squareOp chan<- int) {
// 	sum := 0
// 	digitChan := make(chan int)
// 	go digits(number, digitChan)
// 	for digit := range digitChan {
// 		sum += digit * digit
// 	}
// 	squareOp <- sum
// }

// func calcCubes(number int, cubeOp chan<- int) {
// 	sum := 0
// 	digitChan := make(chan int)
// 	go digits(number, digitChan)
// 	for digit := range digitChan {
// 		sum += digit * digit * digit
// 	}
// 	cubeOp <- sum
// }

// func main() {
// 	number := 589
// 	squareOp, cubeOp := make(chan int), make(chan int)
// 	go calcSquares(number, squareOp)
// 	go calcCubes(number, cubeOp)
// 	fmt.Printf("squares + cubes = %v", <-squareOp+<-cubeOp)
// }

// ----------------------------------------
// func hello(a <-chan bool) {
// 	fmt.Println(<-a)
// }

// func main() {
// 	a := make(chan bool)

// 	// fmt.Println(<-a)

// 	// a <- true

// 	go hello(a)
// 	a <- true

// }

// ----------------------------------------
type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randInt := rand.Intn(999)
		job := Job{i, randInt}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digit %d\n", result.job.id, result.job.randomno, result.sumOfDigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	numberOfWorkers := 10
	createWorkerPool(numberOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}
