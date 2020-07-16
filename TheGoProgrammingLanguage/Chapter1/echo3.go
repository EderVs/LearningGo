// Echo3 using strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// If we don't care about formating.
	// fmt.Println(os.Args[1:])
}
