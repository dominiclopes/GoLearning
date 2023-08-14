package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	call_function_generating_panic()
	fmt.Println("The panic was handled")
}

func call_function_generating_panic() {
	generate_panic()
}

func generate_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered panic:", r)
			fmt.Println("stack trace:")
			debug.PrintStack()

		}
	}()

	panic("I will disrupt everything")
}
