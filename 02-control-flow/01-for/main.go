package main

import "fmt"

func main() {
	sum := 0

	// No parentheses, but braces are always required
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
	// => 10

	// Reset
	sum = 0

	// The above can be write in a more terse way by omitting
	// the initializer and increment statements. Which is
	// essentilly a while loop.
	for sum < 10 {
		sum += 1
	}
	fmt.Println(sum)
	// => 10

	// Without any conditions the loop will run forever
	for {
		sum += 1
		fmt.Println(sum)
	}
}
