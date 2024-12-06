package main

import "testing"

func TestPart1(t *testing.T) {
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
	expected := "41"
	result := Part1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
