package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can also declare methods with pointer receivers.
//
// This means the receiver type has the literal syntax *T for
// some type T. T cannot itself be a pointer like *int, however.
//
// Since methods often need to modify their receiver, pointer
// receivers are more common than value receivers.

// Scale is a method that modifies a Vertex receiver.
// If we remove the * here, the program will not complain
// but the modifications to X and Y will not be performed.
//
// With a value receiver, the Scale method operates on a copy
// of the original Vertex value (the same behavior as with
// other function argument). The Scale method must have a pointer
// receiver to change the Vertex value declared in the main function.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 5

	v.Scale(10)
	fmt.Println(v.Abs()) // 50
}
