package main

import "testing"

func TestPart1(t *testing.T) {
	input := "sample input here"
	expected := "42"
	result := Part1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
