package main

import "fmt"

// Type declaration looks at the right hand side.
//
// var i int
// j := i // int
//
// Untyped numeric constants depend on the precision of the value:
//
// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	v := string("hello")
	fmt.Printf("v is of type %T\n", v)

	v := string("string")
	fmt.Printf("v is of type %T\n", v)
}
