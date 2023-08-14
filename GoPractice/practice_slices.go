package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	var b []int
	fmt.Println(a, b)

	b = append(b, a...)
	fmt.Println(a, b)

	b[3] = 5
	fmt.Println(a, b)

}
