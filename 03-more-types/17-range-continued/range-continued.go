package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)
	// => [0 0 0 0 0 0 0 0 0 0]

	// If you only want the index, you can drop the value from
	// the `for` expression entirely
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(pow)
	// => [1 2 4 8 16 32 64 128 256 512]

	// If you don't want either the index or the value, you can
	// assign to `_`
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	// => 1
	// => 2
	// => 4
	// => 8
	// => 16
	// => 32
	// => 64
	// => 128
	// => 256
	// => 512
}
