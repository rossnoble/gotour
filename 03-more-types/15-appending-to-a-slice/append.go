package main

import (
	"fmt"
)

// It is common to append new elements to a slics, so go
// provides a built-in `append` function. The documentation
// of the built-in package describes `append` using with:
//
// 	 func append(s, []T, vs ...T) []T
//
// The first parameter of `s` of `append` is a slice of type `T`
// and the rest are `T` values to append to the slice.
//
// The resulting value of `append` is a slice containing all
// the elements of the original sluce plus the provided values.
// If the backing array of `s` is too small to fit all the
// given value a bigger array will be allocated. The returned
// slice will point to the newly allocated array.
//
// NOTE: For some reason the `len` and `cap` values printed
// when running this program do not match those in the go tour
// playground example. In that environment, they are returned
// like so:
//
// 	 len=0 cap=0 []
//	 len=1 cap=2 [0]
//	 len=2 cap=2 [0 1]
//	 len=5 cap=8 [0 1 2 3 4]

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)
	// => len=0 cap=0 []

	// You can append to nil slices
	s = append(s, 0)
	printSlice(s)
	// => len=1 cap=1 [0]

	// The slice will grow as needed
	s = append(s, 1)
	printSlice(s)
	// => len=2 cap=2 [0 1]

	// We can add more than one item at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
	// => len=5 cap=6 [0 1 2 3 4 5]
}
