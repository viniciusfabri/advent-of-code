package main

import (
	"fmt"
	"os"
	"testing"
)

// Sample input. Middle numbers ignored
func TestPart2(t *testing.T) {
	b, err := os.ReadFile("sample_input.txt")
	if err != nil {
		fmt.Print(err)
		t.Errorf("Error reading sample_input.txt file")
	}

	input := string(b)

	expected := "4"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Last number ignored
func Test2Part2(t *testing.T) {
	input := `1 2 3 4 10
`

	expected := "1"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test3Part2(t *testing.T) {
	input := `1 2 3 4 10
5 4 0
`

	expected := "2"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test4Part2(t *testing.T) {
	input := `1 2 5 2 7 10
1 2 5 2 4 10
100 97 100
100 97 97 100
`

	expected := "2"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// First number needs to be ignored
func Test5Part2(t *testing.T) {
	input := `1 100 97 94
`

	expected := "1"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
