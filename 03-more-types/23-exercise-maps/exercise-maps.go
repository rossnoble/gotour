package main

import (
	//"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

// Exercise: Maps
//
// Implement `WordCount`. It should returna a map of the counts
// of each "word" in the string`s`. The `wc.Test` function
// runs a test suite against the provided function and prints
// success or failure.
//
// You might find `strings.Fields` helpful.

// f("I am learning Go!") =
// map[string]int{"x":1}
// want:
// map[string]int{"I":1, "am":1, "learning":1, "Go!":1}

// WordCount returns a map of the count of words in a sentence
func WordCount(s string) map[string]int {
	wordList := map[string]int{}
	words := strings.Fields(s)

	for _, w := range words {
		if _, ok := wordList[w]; !ok {
			wordList[w] = 1
		} else {
			wordList[w] = wordList[w] + 1
		}
	}

	return wordList
}

func main() {
	wc.Test(WordCount)
}
