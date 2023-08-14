package main

import "fmt"

func isEmpty(stack *[]int) bool {
	// verify if the stack is empty
	if len(*stack) == 0 {
		return true
	}
	return false

}

func push(stack *[]int, element int) {
	// insert into stack
	fmt.Printf("Pushing element %v to stack %v\n", element, *stack)
	*stack = append(*stack, element)
	fmt.Printf("Elements of the stack after pushing element %v are %v\n", element, *stack)
}

func pop(stack *[]int) {
	// pop out of stack
	if isEmpty(stack) {
		fmt.Println("Stack is empty. Exiting...")
		return
	}
	*stack = (*stack)[:len(*stack)-1]
	fmt.Printf("Elements of the stack after popping the top element is %v \n", *stack)
}

func top(stack *[]int) (int, bool) {
	// find the element at the top of the stack
	if isEmpty(stack) {
		fmt.Println("Stack is empty. Exiting...")
		return 0, false
	}
	top := (*stack)[len(*stack)-1]
	fmt.Printf("Top element of the stack %v is %v.\n", *stack, top)
	return top, true
}

// find if element exists in the stack

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	// Print the top element of the stack
	top(&s)
	// Try to delete the top element of the stack, when the stack is empty
	pop(&s)

	// Insert one element into the stack and verify the top element of the stack
	push(&s, 12)
	fmt.Println(s, len(s), cap(s))
	top(&s)

	// Insert one element into the stack and verify the top element of the stack
	push(&s, 59)
	fmt.Println(s, len(s), cap(s))
	top(&s)

	// Remove the top element of the stack
	pop(&s)
	fmt.Println(s, len(s), cap(s))

	// Insert few elements into the stack and try to pop more than inseted elements from the stack
	push(&s, 100)
	push(&s, 1212)
	push(&s, 9)
	pop(&s)
	pop(&s)
	pop(&s)
	pop(&s)
	pop(&s)
}
