package main

import "fmt"
import "math"

// The above is an alternative to this preferred syntax:
// import(
//   "fmt"
//   "math"
// )

func main() {
	// Note that we are using Printf here instead of Println
	fmt.Printf("You now have %g problems", math.Sqrt(7))
}
