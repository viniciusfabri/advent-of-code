package main

import "testing"

func TestPart2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := "48"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
