package main

import (
	"golang.org/x/tour/pic"
)

// Exercise: Slices
//
// Implement `Pic`. It should return a slice of length `dy`,
// each element of which is a slice of `dx` 8-bit unsigned
// integers. When you run the program, it will display your
// picture, interpreting the integers as grayscale (well,
// bluescale) values.
//
// The choice of image is up to you. Interesting functions
// include `(x+y)/2`, `x*y`, and `x^y`.
//
// Hints: You need to use a loop to allocate each `[]uint8`
// inside the `[][]uint8`. Use uint8(intValue) to convert
// between types.

// Pic returns a slice of length `dy`. Each element is a slice
// of `dx` 8-bit unsigned integers.
func Pic(dx, dy int) [][]uint8 {
	// Build our slice that we will populate
	p := make([][]uint8, dy)

	// For 0 to len(p) / dy, make each element a slice
	// of uint8 items
	for i := range p {
		p[i] = make([]uint8, dx)
	}

	// Iterate over the slice again
	for i := range p {
		// Update each element inside the nested slices
		for j := range p[i] {
			p[i][j] = uint8(0) + uint8(i)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
