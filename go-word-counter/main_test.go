package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	buffer := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4
	result := count(buffer)

	if result != expected {
		t.Errorf("Expected %d but result is %d", expected, result)
	}
}
