package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart2WithSampleInput(t *testing.T) {
	b, err := os.ReadFile("sample_input.txt")
	if err != nil {
		fmt.Print(err)
		t.Errorf("Error reading sample_input.txt file")
	}

	str := string(b)

	expected := 31
	result := Part2(str)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2WithRealInput(t *testing.T) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		t.Errorf("Error reading input.txt file")
	}

	str := string(b)

	expected := 21024792
	result := Part2(str)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
