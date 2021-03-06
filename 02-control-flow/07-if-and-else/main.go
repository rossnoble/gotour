package main

import (
	"fmt"
	"math"
)

func pow(v, n, lim float64) float64 {
	if v := math.Pow(v, n); v < lim {
		return v
	} else {
		// v is accessible in here as well
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 28),
	)
}
