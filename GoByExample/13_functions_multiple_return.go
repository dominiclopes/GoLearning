package main

import "fmt"

func vars() (int, int) {
	return 2, 3
}

func main() {
	a, b := vars()
	fmt.Println(a, b)

	_, c := vars()
	fmt.Println(c)
}
