package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (vert Vertex) Abs() float64 {
	return math.Sqrt(vert.X*vert.X + vert.Y*vert.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// => 5
}
