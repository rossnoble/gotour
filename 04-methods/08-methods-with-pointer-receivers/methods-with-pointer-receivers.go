package main

import (
	"fmt"
	"math"
)

// There are two reasons to use a pointer receiver:
//
// 1) So that the method can modify the value that its
// receiver points to.
//
// 2) To avoid copy the value on each method call. This
// can be more effienct if the receiver is a large
// struct, for example.

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here we can see how, like the the Scale function, the
// and Abs function receives a type Vertex even though
// the Abs does *not* need to modify its receiver.
//
// In general, all methods on a given type should have
// either value or pointer receivers, but not a mixture
// of both.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}

	// => Before scaling: &{X:3 Y:4}, Abs: 5
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)

	// => After scaling: &{X:15 Y:20}, Abs: 25
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
