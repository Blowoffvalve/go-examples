package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the point
	*p = 21         // set i through the point
	fmt.Println(i)  // set the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the point
	fmt.Println(j)
}
