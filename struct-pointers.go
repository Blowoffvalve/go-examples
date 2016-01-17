package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// Struct fields can be accessed through a struct point, the indirection is
	// transparent
	p.X = 1e9
	fmt.Println(v)
}
