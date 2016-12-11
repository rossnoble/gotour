package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs has been converted to a regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale has been converted to a regular function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v)) // 5

	// Functions with have a pointer as an argument
	// much be passed a pointer
	Scale(&v, 10)

	// Works the same as before
	fmt.Println(Abs(v)) // 50
}
