package main

import "fmt"

type Vertex struct {
	// We can use the same type assignment as with
	// regular variables
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 2}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{2, 4} // Type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
	// => {1 2}
	// => {2 0}
	// => {0 0}
	// => &{2 4}
}
