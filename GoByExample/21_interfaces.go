package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2 * r.length * r.breadth
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("%v is of type %T\n", g, g)
	fmt.Printf("Area of %v is %v\n", g, g.area())
	fmt.Printf("Perimeter of %v is %v\n", g, g.perimeter())
}

func main() {
	r := rect{2, 3}
	measure(r)

	c := circle{4}
	measure(c)
}
