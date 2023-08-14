package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
	fmt.Println(ival, &ival)

}

func zeroPtr(iptr *int) {
	*iptr = 0
	fmt.Println(*iptr, iptr)

}

func main() {
	i := 1
	fmt.Println(i, &i)

	zeroVal(i)
	fmt.Println(i, &i)

	zeroPtr(&i)
	fmt.Println(i, &i)
}
