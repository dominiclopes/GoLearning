package main

import "fmt"

func sum(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("sum =", sum)
}

func main() {
	sum(1, 2)
	sum(2, 4, 6)

	c := []int{5, 6, 7}
	sum(c...)
}
