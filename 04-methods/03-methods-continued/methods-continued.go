package main

import (
	"fmt"
	"math"
)

// You can also define methods on non-struct types. Here we
// have a simple float that we can now add methods to.
type MyFloat float64

// You can only declare a method with a receiver whose type
// is defined in the same package as the method. You cannot
// declare a method with a receiver whose type is defined in
// another package (which includes built-in types like int).

// Abs is a method that can be called on
// our float type defined above
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	//  MyFloat is being to perform type conversion here
	f := MyFloat(-math.Sqrt2)

	// This could be written like this instead
	var n MyFloat = -math.Sqrt2

	// => 1.4142135623730951
	fmt.Println(f.Abs())

	// => true
	fmt.Println("Equal?", f == n)
}
