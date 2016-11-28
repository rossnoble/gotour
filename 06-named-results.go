package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// Naked return
	// Equivalent to: return x, y
	return
}

func main() {
	fmt.Println(split(17))
}
