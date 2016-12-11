package main

import (
	"fmt"
	"math"
)

// Function are values too. They can be passed around just like
// other values. Function value may be used as function arguments
// and return values.

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// Use passed values
	fmt.Println(hypot(5, 12))
	// => 13

	// Use default values
	fmt.Println(compute(hypot))
	// => 5

	// Pass in any function that matches the signature
	fmt.Println(compute(math.Pow))
	// => 81

	// This would error because it doesn't match
	// fmt.Println(compute(math.Sqrt))
}
