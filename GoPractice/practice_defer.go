package main

import (
	"fmt"
	"sync"
)

// type Vertex struct {
// 	X int
// 	Y int
// }

// func (v *Vertex) SumOfStructIntergers() {
// 	fmt.Println("Getting the sum of intergers")
// 	fmt.Println("The sum of integers is:", v.X+v.Y)

// }

// func Square(num int) {
// 	fmt.Println("Getting the square of num:", num)
// 	fmt.Println("The square of num:", num, "is:", num*num)
// }

// func main() {
// 	fmt.Println("This is main function - output1")

// 	// Call a function deferring functions within
// 	show_level1()

// 	// Defer a function
// 	defer Square(2)

// 	// Defer a method
// 	v := Vertex{}
// 	defer v.SumOfStructIntergers()

// 	// Defer an anonymous function
// 	defer func(num float64) {
// 		fmt.Println("Getting the absoulute value of num:", num)
// 		fmt.Println("The absoulute value of num:", num, "is:", math.Abs(num))
// 	}(-10.12)

// 	// Defer a function to verify the parameters are evaluated before
// 	defer func(greeting string) {
// 		fmt.Println("The greeting says:", greeting)
// 	}(func() (greeting string) {
// 		greeting = "Hey there!"
// 		fmt.Println("I will greet saying:", greeting)
// 		return
// 	}())
// 	fmt.Println("This is main function - output2")
// }

// func show_level1() {
// 	defer fmt.Println("Shown - Level 1 - 1")
// 	fmt.Println("Shown - Level 1 - 2")
// 	defer fmt.Println("Shown - Level 1 - 3")
// 	show_level2()
// }

// func show_level2() {
// 	defer fmt.Println("Shown - Level 2 - 1")
// 	fmt.Println("Shown - Level 2 - 2")
// 	defer fmt.Println("Shown - Level 2 - 3")
// }

// ---------------------------
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
}
