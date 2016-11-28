package main

import "fmt"

// This is not allowed
k := "some words"

func main() {
	var i, j int = 1, 2

	// k will be declared as an int
  // := is only available inside the function scope
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, c, k, python, java)
	// => 1, 2, 3, true, false, "no!"
}
