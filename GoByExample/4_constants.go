package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	// constants
	fmt.Printf("s = %v, %T\n", s, s)
	const n = 500000000
	fmt.Printf("n = %v, %T\n", n, n)
	const d = 3e20 / n
	fmt.Printf("d = %v, %T\n", d, d)
	fmt.Println(int(d), int64(d))
	fmt.Println(math.Sin(n))
}
