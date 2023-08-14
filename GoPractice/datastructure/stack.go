package main

import (
	"errors"
	"fmt"
)

type stack []string

var errEmptyStack = errors.New("stack is empty")

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Head() (string, error) {
	if s.IsEmpty() {
		return "", errEmptyStack
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errEmptyStack
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, nil
}

func main() {
	var myStack stack
	fmt.Println(myStack)

	is_empty := myStack.IsEmpty()
	fmt.Printf("My stack %v is_empty: %v\n", myStack, is_empty)

	str, err := myStack.Head()
	fmt.Printf("The head of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Pop()
	fmt.Printf("The popped elemtn of my stack %v is %q. Error: %v\n", myStack, str, err)

	myStack.Push("This")
	myStack.Push("Is")
	myStack.Push("Amazing")
	fmt.Println(myStack)

	is_empty = myStack.IsEmpty()
	fmt.Printf("My stack %v is_empty: %v\n", myStack, is_empty)

	str, err = myStack.Head()
	fmt.Printf("The head of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Pop()
	fmt.Printf("The popped elemtn of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Head()
	fmt.Printf("The head of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Pop()
	fmt.Printf("The popped elemtn of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Head()
	fmt.Printf("The head of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Pop()
	fmt.Printf("The popped elemtn of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Head()
	fmt.Printf("The head of my stack %v is %q. Error: %v\n", myStack, str, err)

	str, err = myStack.Pop()
	fmt.Printf("The popped elemtn of my stack %v is %q. Error: %v\n", myStack, str, err)

}
