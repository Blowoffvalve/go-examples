package main

import "fmt"

// See http://blog.golang.org/defer-panic-and-recover for more detail
func main() {
	// Defer the execution of a functions until the surrounding function returns
	// Arguments is evaluated immediately
	defer fmt.Println("world")

	// Defer statements execute in last-in-first-out order.
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}
