package main

import "fmt"

func main() {
	// for loops
	i := 1
	for i <= 3 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 1; j < 5; j++ {
		fmt.Println("j = ", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println("k = ", k)
	}
}
