package main

import "fmt"

func merge(a []int, b []int) (final []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	if i < len(a) {
		final = append(final, a[i:]...)
	}

	if j < len(b) {
		final = append(final, b[j:]...)
	}
	fmt.Printf("%v + %v -> %v\n", a, b, final)
	return
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

// func main() {
// 	sliceToSort := []int{12, 34, 1, 5, 7, 23}
// 	mergeSort(sliceToSort)
// }
