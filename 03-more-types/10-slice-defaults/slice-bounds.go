package main

import (
	"fmt"
)

// When slicing, you may omit the high or low bounds to use
// their defaults instead. The default for the low bound is
// zero, and the default for the high bound is the length
// of the slice. These slice expressions are equivalent:
//
//  a[0:10]
//  a[:10]
//  a[0:]
//  a[:]

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// Update s to slice of original slice
	s = s[1:4]
	fmt.Println(s)
	// => [3 5 7]

	// Update s again to slice of sliced slice
	s = s[:2]
	fmt.Println(s)
	// => [3 5]

	s = s[1:]
	fmt.Println(s)
	// => [5]
}
