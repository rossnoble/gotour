package main

import (
	"fmt"
	"math"
)

// Type conversion examples T(v)
//
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
//
// If we want to change a variable's type, we must explicitly
// explicitly perform a conversion. This would throw an error:
//
// var u int = f

func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)
	// => 42 42 42

	var x, y int = 3, 4
	var f64 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f64)
	fmt.Println(x, y, z)
	// => 3 4 5
}
