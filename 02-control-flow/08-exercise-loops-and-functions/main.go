package main

import (
	"fmt"
	"math"
)

const Diff = 0.00001

func Sqrt(x float64) (float64, int) {
	z := 1.0
	timesRun := 0

	for {
		next := z - (z*z-x)/(2*z)

		_ = "breakpoint"
		if d := math.Abs(next - z); d < Diff {
			// Close enough
			return next, timesRun
		}

		timesRun += 1

		// Try next one
		z = next
	}

	return z, timesRun
}

func main() {
	var testValue float64 = 9988877665544
	expected, t := Sqrt(testValue)
	actual := math.Sqrt(testValue)
	delta := math.Abs(actual - expected)

	fmt.Println("expected:", expected)
	fmt.Println("actual:", actual)
	fmt.Println("diff:", delta)
	fmt.Println("times run:", t)
}
