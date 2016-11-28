package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Shorthand type declaration of the above.
// Both x and y are integers here.
func addAgain(x, y int) int {
	return x + y
}

func main() {
	// This will work becasuse we are passing two integers
	fmt.Println(add(42, 13))

	// This will work becasuse we are passing two integers
	fmt.Println(addAgain(42, 13))

	// This will fail because we are passing a string
	fmt.Println(add(42, "13"))
}
