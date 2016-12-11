package main

import (
	"fmt"
)

func main() {
	// Arrays are defined using this notation: [n]T.
	// An array's length is part of its type so it
	// cannot be resized.

	// Define a new string array of length 2
	var a [2]string
	// Add an item at index 0
	a[0] = "Hello"
	// Add an item at index 1
	a[1] = "World"

	fmt.Println(a[0], a[1])
	// => Hello World
	fmt.Println(a)
	// => [Hello World]

	// Populate the array at build time
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// => [2, 3, 5, 7, 11, 13]

	// Adding 7 items would not work, however:
	// primes := [6]int{2, 3, 5, 7, 11, 13, 17}
}
