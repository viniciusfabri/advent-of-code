package main

import (
	"testing"
)

func TestPart2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	expected := "6"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
