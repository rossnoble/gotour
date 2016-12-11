package main

import (
	"fmt"
)

// A Vertex is a struct
type Vertex struct {
	Lat, Long float64
}

// If the top-level type is just a type name, you can omit it
// from the elements of the literal.
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  // `Vertex` not needed
	"Google":    {37.42202, -122.08408}, // `Vertex` not needed
}

func main() {
	fmt.Println(m)
	// => map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
