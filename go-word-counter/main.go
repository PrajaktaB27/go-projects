package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin)) //std input
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wordcount := 0
	for scanner.Scan() {
		wordcount++
	}

	return wordcount
}
