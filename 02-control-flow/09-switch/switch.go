package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Note that Print omits the newline char
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// Name of os
		fmt.Println("Linux.")
	}
}
