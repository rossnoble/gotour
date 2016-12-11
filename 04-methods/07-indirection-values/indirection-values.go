package main

import (
	"fmt"
	"math"
)

// Here we have the same equivalent setup as in
// the previous exercise to show how the equivalent
// thing happens in the reverse direction.
//
// Functions that take a value argument must take a
// value of that specific type. For example:
//
//	 var v Vertex
//	 fmt.Println(AbsFunc(v))  // OK
//	 fmt.Println(AbsFunc(&v)) // Compile error!
//
// In other words, fmt.Println complains when
// passed a pointer.
//
// Meanwhile methods with value receivers take either a
// value or a pointer as the receiver when they are called:
//
//	 var v Vertex
//	 fmt.Println(v.Abs()) // OK
//	 p := &v
//	 fmt.Println(p.Abs()) // OK
//
// In this case, the method call p.Abs() is interpreted as (*p).Abs().

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	// These behave as expected, both printing 5
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	// We can achieve the same thing with a pointer
	p := &Vertex{4, 3}

	// This is interpreted as (*p).Abs() and will print 5
	fmt.Println(p.Abs())

	// This will also print 5
	fmt.Println(AbsFunc(*p))
}
