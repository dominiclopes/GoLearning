package main

import (
	"fmt"
)

// -------------------------------
// func main() {
// 	select {}
// }

// -------------------------------
// func main() {
// 	var ch chan bool
// 	select {
// 	case <-ch:
// 	}
// }

// -------------------------------
// func main() {
// 	var ch chan bool
// 	select {
// 	case <-ch:
// 	default:
// 		fmt.Println("default select case")
// 	}
// }

// -------------------------------
func server1(ch chan<- string) {
	// time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan<- string) {
	// time.Sleep(3 * time.Second)
	ch <- "from server2"
}
func main() {
	output1, output2 := make(chan string), make(chan string)
	go server2(output2)
	go server1(output1)

	// time.Sleep(6 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
		// default:
		// 	fmt.Println("default select case")
	}
}
