package main

import (
	"testing"
)

func TestPart2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	expected := "9"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func Test2Part2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	expected := "9"
	result := Part2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
