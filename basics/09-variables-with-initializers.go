package main

import "fmt"

// This will assign 2 and 4 to a and b respectively
var a, b int = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	b = 3

	fmt.Println(a, b, c, python, java)
	// => 1 3 true false "no!"
}
