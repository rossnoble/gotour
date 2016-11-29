package main

import "fmt"

// Constants can be character, string, boolean, or numeric values.
// They cannot be declared using the := syntax.
const Pi = 3.14

// You can't redefine or modify constants. These are no good:
// const Pi = 1
// var Pi = 1
// Pi = 1

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	// => Hello 世界

	fmt.Println("Happy", Pi, "Day")
	// => Happy 3.14 Day

	const Truth = true
	fmt.Println("Go rules?", Truth)
	// => Go rules?
}
