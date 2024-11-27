package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//if 1 count lines instead of words
	lines := flag.Bool("1", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines)) //std input
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	//if flag not set, count words
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	wordcount := 0
	for scanner.Scan() {
		wordcount++
	}

	return wordcount
}
