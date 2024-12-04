package main

import (
	"fmt"
	"strings"
)

// Exported function to be called by the main application
func Part2(input string) string {
	input = strings.Trim(input, "\n")

	grid := strings.Split(input, "\n")

	count := countXMASPart2(grid)

	return fmt.Sprint(count)
}

func countXMASPart2(grid []string) int {
	rows := len(grid)
	// Not enough rows for an X-MAS
	if rows < 3 {
		return 0
	}

	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			if isXMAS(grid, r, c) {
				count++
			}
		}
	}

	return count
}

func isXMAS(grid []string, r, c int) bool {
	// It's only an X-MAS when the center character is "A"
	if grid[r][c] != 'A' {
		return false
	}

	// Ensure we can check the other characters (as in they exist in the grid)
	if r-1 < 0 || r+1 >= len(grid) || c-1 < 0 || c+1 >= len(grid[r]) {
		return false
	}

	upperLeftCharacter := grid[r-1][c-1]
	upperRightCharacter := grid[r-1][c+1]
	lowerLeftCharacter := grid[r+1][c-1]
	lowerRightCharacter := grid[r+1][c+1]

	// Check if we formed a X-MAS
	if upperLeftCharacter == 'M' && lowerRightCharacter == 'S' {
		if upperRightCharacter == 'M' && lowerLeftCharacter == 'S' {
			return true
		}

		if upperRightCharacter == 'S' && lowerLeftCharacter == 'M' {
			return true
		}
	}

	if upperLeftCharacter == 'S' && lowerRightCharacter == 'M' {
		if upperRightCharacter == 'M' && lowerLeftCharacter == 'S' {
			return true
		}

		if upperRightCharacter == 'S' && lowerLeftCharacter == 'M' {
			return true
		}
	}

	return false
}
