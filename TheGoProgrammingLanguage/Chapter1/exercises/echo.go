// Echo: Exercises 1.1 and 1.2
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Exercise 1.1
	fmt.Println("Exercise 1.1")
	fmt.Println(strings.Join(os.Args, " "))
	// Exercise 1.2
	fmt.Println("Exercise 1.2")
	sep := " "
	for i, arg := range os.Args {
		s := strconv.Itoa(i) + sep + arg
		fmt.Println(s)
	}
}
