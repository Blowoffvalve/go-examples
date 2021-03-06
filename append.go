package main

import "fmt"

// see https://blog.golang.org/go-slices-usage-and-internals for more details
func main() {
	var a []int
	printSlice("a", a)

	// append works on nil slices
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
