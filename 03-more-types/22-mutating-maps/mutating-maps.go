package main

import (
	"fmt"
)

// You can mutate a map in many ways. Here are some examples:
//
// Insert or update an element in map `m`:
//
//   m[key] = elem
//
// Retreive an element:
//
//   elem = m[key]
//
// Delete an element:
//
//   delete(m, key)
//
// Test that a key is prsent with a two-value assignment:
//
//   elem, ok = m[key]
//
// If `key` in `m`, `ok` is `true`, otherwise `ok` is false. If
// `key` is not in the map, then `elem` is the zero value for the
// map's element type.
//
// Note: if `elem` or `ok` have not yet been declared, you can
// use the short declaration form:
//
//   elem, ok := m[key]

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	// => The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	// => The value: 48

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// => The value: 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	// => The value: 0 Present? false
}
