// Dup1 prints the count and the text of each line that appears more than
// once in the input. It reads from stin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsPerFile := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsPerFile, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsPerFile, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("Word:")
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println("Files:")
			for filename, m := range countsPerFile[line] {
				fmt.Printf("%d\t%s\n", m, filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsPerFile map[string]map[string]int, filename string) { // Reference pass
	input := bufio.NewScanner(f)
	var line string
	for input.Scan() {
		line = input.Text()
		if counts[line] == 0 {
			countsPerFile[line] = make(map[string]int)
		}
		counts[line]++
		countsPerFile[line][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
