package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// this is wrong?  No tuples?
	// the range behavior should be consistent. If we just want the index, it
	// should be `for i, _ := range pow`. `for i := range pow` i should contain
	// (key, value)
	for i, _ := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
