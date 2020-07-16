// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// for loop is the only loop statement in Go.
	// There are no parethesis in the trhee components of a for loop.
	for i := 1; i < len(os.Args); i++ { // i++ is a statement, so j = i++ is ilegal.
		s += sep + os.Args[i] // Concat strings.
		sep = " "
	}
	// a traditional "while" loop is using for without initialization and post.
	for false {
		// ...
	}
	// We can also add infinite loops like this:
	for {
		break
	}
	fmt.Println(s)
}
