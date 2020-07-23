// Fetchall fetches URLs in parallel and reposrts their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // channel
	for i, url := range os.Args[1:] {
		go fetch(url, strconv.Itoa(i)+".txt", ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, destinationFilename string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	destination, err := os.Create(destinationFilename)
	if err != nil {
		ch <- fmt.Sprintf("while creating %s: %v", destinationFilename, err)
		return
	}
	defer destination.Close()
	nbytes, err := io.Copy(destination, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
