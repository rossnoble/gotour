package main

import (
	"fmt"
)

// The zero value of a slice is nil. A nil slice has a length
// and a capacity of 0 and no underlying array.

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	// => [] 0 0

	// This will be true
	if s == nil {
		fmt.Println("nil!")
		// => nil!
	}
}
