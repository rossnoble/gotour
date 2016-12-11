package main

import (
	"fmt"
)

// An array has a fixed size. A slice, however, has a
// dynamically sized, flexible view into the elements
// of an array. In practice, slices are more common
// than arrays. They are defined with the notation: []T.

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Create a slice from primees
	var s []int = primes[1:4]
	// Create with inferred type
	k := primes[0:5]

	fmt.Println(s)
	// => [3, 5, 7]
	fmt.Println(k)
	// => [2, 3, 5, 7, 11]
}
