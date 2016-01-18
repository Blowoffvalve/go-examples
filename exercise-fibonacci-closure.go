package main

import "fmt"

// fibonacci is a function that returns a function that returns an int8
func fibonacci() func() int {
	first, second := 1, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
