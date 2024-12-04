package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, err := os.ReadFile("sample_input.txt")
	if err != nil {
		fmt.Print(err)
		t.Errorf("Error reading sample_input.txt file")
	}

	input := string(b)

	expected := "2"
	result := Part1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart1WithRealInput(t *testing.T) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		t.Errorf("Error reading input.txt file")
	}

	str := string(b)

	expected := "252"
	result := Part1(str)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
