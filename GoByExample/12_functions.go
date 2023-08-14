package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Functions
	res1 := plus(1, 2)
	fmt.Println(res1)

	res2 := plusplus(1, 2, 3)
	fmt.Println(res2)
}
