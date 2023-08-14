package main

import (
	"fmt"

	greetings "location_to_greetings"
)

func main() {
	message, err := greetings.Hello("Dominic")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(message)

	messages, err := greetings.Hellos([]string{"Anuja", "Niharika", "Rahul"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(messages)
}
