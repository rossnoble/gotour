package main

import (
	"fmt"
	"math"
)

func main() {
	// Pi must be capitalized to be exported from the
	// math package. It is capitalized here as well.
	fmt.Println(math.Pi)
}
