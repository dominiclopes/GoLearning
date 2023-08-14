package greetings

import "fmt"

func init() {
	fmt.Println("Inside the init of greetings_2.go")
}

func GoodMorrow(name string) (message string) {
	message = fmt.Sprintf("Hi %v. Good Morrow!", name)
	return
}
