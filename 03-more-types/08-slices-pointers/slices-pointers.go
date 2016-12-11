package main

import (
	"fmt"
)

// A slice does not store any data. It just describes
// a section of an underlying array. Changing the elements
// of a slice changes the corresponding elements of its
// underlying array. Other slices of the same array will
// also see those changes.

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	// => [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	// => [John, Paul] [Paul, George]

	b[0] = "XXX" // Paul => XXX
	fmt.Println(a, b)
	// => [John, XXX] [XXX, George]
	fmt.Println(names)
	// => [John XXX George Ringo]
}
