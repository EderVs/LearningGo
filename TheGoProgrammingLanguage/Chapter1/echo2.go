// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// Different types of declarations:
	// s := ""
	// var s string
	// var s = "" // This is used when declaring multiple variables.
	// var s string = "" // This is redundant.

	// Like a for loop in python.
	// Produces a pair of values: the index and the value of the element at that index.
	// _ is the black identifier.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
