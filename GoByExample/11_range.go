package main

import "fmt"

func main() {
	// range
	numSlice := []int{1, 2, 3}
	for _, v := range numSlice {
		fmt.Println(v)
	}

	numArray := [3]string{"this", "is", "me"}
	for i, v := range numArray {
		fmt.Println(i, "->", v)
	}

	kvs := map[string]bool{"hello": true, "are": false, "you": true}
	for k, v := range kvs {
		fmt.Println(k, "=", v)
	}

	for i, v := range "go" {
		fmt.Println(i, "-", v)
	}
}
