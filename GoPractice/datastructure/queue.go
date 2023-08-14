package main

import (
	"errors"
	"fmt"
)

type Queue []int

var ErrEmptyQueue = errors.New("empty queue")

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(element int) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, nil
}

func main() {
	var myQueue Queue
	fmt.Println(myQueue)

	is_empty := myQueue.IsEmpty()
	fmt.Printf("My Queue %v is empty: %v\n", myQueue, is_empty)

	element, err := myQueue.Dequeue()
	fmt.Printf("The first element in my queue %v is %v. Error: %v\n", myQueue, element, err)

	myQueue.Enqueue(34)
	fmt.Println(myQueue)

	myQueue.Enqueue(123)
	fmt.Println(myQueue)

	is_empty = myQueue.IsEmpty()
	fmt.Printf("My Queue %v is empty: %v\n", myQueue, is_empty)

	element, err = myQueue.Dequeue()
	fmt.Printf("The first element in my queue %v is %v. Error: %v\n", myQueue, element, err)

	is_empty = myQueue.IsEmpty()
	fmt.Printf("My Queue %v is empty: %v\n", myQueue, is_empty)

	element, err = myQueue.Dequeue()
	fmt.Printf("The first element in my queue %v is %v. Error: %v\n", myQueue, element, err)

	is_empty = myQueue.IsEmpty()
	fmt.Printf("My Queue %v is empty: %v\n", myQueue, is_empty)

	element, err = myQueue.Dequeue()
	fmt.Printf("The first element in my queue %v is %v. Error: %v\n", myQueue, element, err)

}
