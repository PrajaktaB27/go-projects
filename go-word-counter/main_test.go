package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	buffer := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4
	result := count(buffer, false, false)

	if result != expected {
		t.Errorf("Expected %d but result is %d", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	buffer := bytes.NewBufferString("word1 word2 word3 word4\n line2\nline3 word5")

	expected := 3
	result := count(buffer, true, false)

	if result != expected {
		t.Errorf("Expected %d but result is %d", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	buffer := bytes.NewBufferString("this is to count bytes")

	expected := len("this is to count bytes")
	result := count(buffer, false, true)

	if result != expected {
		t.Errorf("Expected %d but result is %d", expected, result)
	}
}
