package main

import (
	"fmt"
)

// A map maps keys to values. The zero value of a map is `nil`.
// A `nil` map has no keys, nor can keys be added. The `make`
// function returns a map of the given type, initialized and
// ready for use.

// A Vertex is a struct of floats
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	// => {40.68433 -74.39967}
}
