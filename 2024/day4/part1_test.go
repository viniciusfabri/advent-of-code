package main

import "testing"

func TestPart1(t *testing.T) {
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
	expected := "18"
	result := Part1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
