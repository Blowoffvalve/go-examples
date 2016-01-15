package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 1e-10
	for math.Abs((z*z-x)/x) > delta {
		z = z - (z*z-x)/2/z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt2 - Sqrt(2))
}
