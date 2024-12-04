package main

import (
	"fmt"
	"strings"
)

// Exported function to be called by the main application
func Part1(input string) string {
	// Split the input into rows
	grid := strings.Split(input, "\n")

	// Count the number of occurrences of the word "XMAS"
	count := countXMAS(grid)

	return fmt.Sprint(count)
}

func countXMAS(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		fmt.Println("The grid is empty.")
		return 0
	}

	cols := 0
	for _, row := range grid {
		if len(row) > cols {
			cols = len(row)
		}
	}

	word := "XMAS"
	wordLen := len(word)
	directions := [][]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Down-right
		{1, -1},  // Down-left
		{-1, 1},  // Up-right
		{-1, -1}, // Up-left
	}

	isMatch := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLen; i++ {
			nr, nc := r+dr*i, c+dc*i
			// Check boundaries
			if nr < 0 || nr >= rows || nc < 0 || nc >= len(grid[nr]) || grid[nr][nc] != word[i] {
				return false
			}
		}
		return true
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < len(grid[r]); c++ {
			for _, dir := range directions {
				if isMatch(r, c, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}
