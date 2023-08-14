package main

import "fmt"

type rect struct {
	length  int
	breadth int
}

func (r rect) area() int {
	return r.length * r.breadth
}

func (r *rect) perimeter() int {
	return 2 * r.length * r.breadth
}

func main() {
	r := rect{length: 10, breadth: 5}
	fmt.Println(r)
	fmt.Println("area =", r.area(), "perimeter =", r.perimeter())

	rp := &r
	fmt.Println(rp, *rp)
	fmt.Println("area =", rp.area(), "perimeter =", rp.perimeter())
	fmt.Println("area =", (*rp).area(), "perimeter =", (*rp).perimeter())

}
