package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i

	// Read i through the p pointer
	fmt.Println(*p) // => 42

	// Update i through the p pointer
	*p = 21

	fmt.Println(i) // => 21

	// Now point p to j instead
	p = &j

	// Divide j through the p pointer
	*p = *p / 37

	fmt.Println(j)
	// => 42
	// => 21
	// => 47
}
