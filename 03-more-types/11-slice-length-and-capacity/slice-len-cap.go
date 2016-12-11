package main

import (
	"fmt"
)

// A slice has both a length and a capacity. The length of a
// slice is the number of elements it contains. The capacity
// of a slice is the number of elements in the underlying
// array, counting from the first element in the slice.
//
// The length and capacity of a slice `s` can be obtained
// using the expressions `len(s)` and `cap(s)`.
//
// You can extend a slice's length by reslicing it, providing
// it has sufficient capacity.

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// => len=6 cap=6 [2 3 5 7 11 13]

	// SLice the slice to give is a length of zero
	s = s[:0]
	printSlice(s)
	// => len=0 cap=6 []

	// Extend the length again
	s = s[:4]
	printSlice(s)
	// => len=4 cap=6 [2 3 5 7]

	// Drop the first two values
	s = s[2:]
	printSlice(s)
	// => len=2 cap=6 [5 7]

	// Extend it again
	s = s[1:4]
	printSlice(s)
	// => len=3 cap=6 [7 11 13]

	// Extending the slice beyond the capacity of the
	// underlying array will cause an error however:
	// s = s[:4]
	// printSlice(s)
	// => panic
}
