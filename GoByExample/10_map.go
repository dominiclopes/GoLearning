package main

import "fmt"

func main() {
	// maps
	m := make(map[string]int)
	fmt.Println(m)
	m["k1"] = 7
	m["k2"] = 8
	fmt.Println(m, len(m))

	delete(m, "k2")
	fmt.Println(m)
	fmt.Println(m["k2"])

	v, prs := m["k2"]
	fmt.Println(v, prs)

	v, prs = m["k1"]
	fmt.Println(v, prs)

	n := map[string]int{"a": 1, "b": 2}
	fmt.Println(n)

	o := map[string]int{
		"day1": 123,
		"day2": 223,
	}
	fmt.Println(o)
}
