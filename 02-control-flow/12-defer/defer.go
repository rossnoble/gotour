package main

import "fmt"

func main() {
	count()
}

func simpleEx() {
	// This will not be called until the surrounding function
	// returns. It is evaluated immediately however.
	defer fmt.Println("world")

	fmt.Println("hello")
}

func count() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
