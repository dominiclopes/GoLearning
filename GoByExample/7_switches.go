package main

import (
	"fmt"
	"time"
)

func main() {
	// basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// can use commas to separate multiple conditions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its the weekday")
	}

	// switch without an expression is an alternate way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon")
	default:
		fmt.Println("It's after Noon")
	}

	// type swith compares types instead of values
	// used to discover type of an interface value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an int")
		case bool:
			fmt.Println("I'm a boolean")
		default:
			fmt.Printf("The type is %T\n", t)
		}
	}
	whatAmI(8)
	whatAmI("hello")
	whatAmI(true)
}
