package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// Scale method of a vertex
// Methods with pointer receivers take either a value
// or a pointer as the receiver when they are called:
//
// For example:
//
// 		var v Vertex
// 		v.Scale(5)  // OK
// 		p := &v
// 		p.Scale(10) // OK
//
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale function that scales a pointer
// Functions with a pointer argument must take a pointer
//
// For example:
//
// 		ScaleFunc(v)  // Compile error!
// 		ScaleFunc(&v) // OK
//
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	// v becomes {6, 8} as expected
	v.Scale(2)

	// Note how we have to pass v as a pointer here
	// v becomes {60, 80}
	ScaleFunc(&v, 10)

	// p is a pointer to a vertex
	p := &Vertex{4, 3}

	// p is a pointer so we can call the method
	// p becomes &{12, 9}
	p.Scale(3)

	// ScaleFunc will accept a pointer
	// p becomes &{120, 90}
	ScaleFunc(p, 10)

	fmt.Println(v, p)
	// => {60, 80} &{120, 90}
}
