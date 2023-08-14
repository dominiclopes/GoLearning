package main

import "fmt"

func bubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		fmt.Println("i", i)
		for j := 0; j < len(a)-1-i; j++ {
			fmt.Println("j", j)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
	return a
}

func main() {
	sliceToSort := []int{12, 34, 1, 5, 7, 23}
	bubbleSort(sliceToSort)
}
