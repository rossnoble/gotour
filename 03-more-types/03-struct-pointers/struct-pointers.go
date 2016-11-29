package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{4, 6}

	// Pointer for a specific field
	p := &v.X

	// Update through pointer
	*p = 8

	fmt.Println(v)
	// => {8, 6}

	v2 := Vertex{10, 1}

	// Pointer for the entire struct
	p2 := &v2

	// Note how the * is omitted here
	p2.Y = 5

	fmt.Println(v2)
	// => {10, 5}
}
