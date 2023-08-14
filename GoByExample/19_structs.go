package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"dominic", 25})
	fmt.Println(person{name: "Dareen", age: 22})
	fmt.Println(person{name: "Janice"})
	fmt.Println(&person{"Ann", 10})

	// its idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Lorson"))

	s := person{
		name: "Superman",
		age:  1000232,
	}
	fmt.Println(s)

	sp := &s
	(*sp).age = 12
	fmt.Println(s)

	sp.age = 121
	fmt.Println(s)

}
