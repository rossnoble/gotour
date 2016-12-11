package main

import (
	"fmt"
	"math"
)

// Here we are rewriting example 01 such that
// Abs is a regular function instead of a method

type Vertex struct {
	X, Y float64
}

// Abs is a function that accepts a vertex and
// returns its square root.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(Abs(v)) // 5
}
