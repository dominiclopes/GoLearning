package main

import "fmt"

func main() {
	// slices
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println("len =", len(s), "cap =", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set s =", s)
	fmt.Println("len =", len(s), "cap =", cap(s))

	s = append(s, "d")
	fmt.Println("set s =", s)
	fmt.Println("len =", len(s), "cap =", cap(s))

	s = append(s, "e", "f")
	fmt.Println("set s =", s)
	fmt.Println("len =", len(s), "cap =", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("set c =", c)
	fmt.Println("len =", len(c), "cap =", cap(c))

	fmt.Println(s[:5])
	fmt.Println(s[2:])

	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println(len(twoD), cap(twoD))
	fmt.Printf("type of twoD = %T\n", twoD)
	for i := 0; i < 3; i++ {
		fmt.Printf("type of twoD[%v] = %T\n", i, twoD[i])
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
