package main

import (
	"fmt"

	day6 "github.com/viniciusfabri/aoc-2024/2024/day6/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	roomGrid := day6.ParseInputIntoGrid(input)

	// We start with 1 because we include the guard starting position
	distinctPositionsVisited := 1
	for {
		guardRow, guardColumn, direction := day6.FindSecurityGuardPositionAndDirection(roomGrid)

		if day6.GuardCanLeave(roomGrid, guardRow, guardColumn, direction) {
			break
		}

		if day6.HasObstaculeAhead(roomGrid, guardRow, guardColumn, direction) {
			direction = day6.NextDirection(direction)
			// Update room grid with guard facing new direction and restart loop
			roomGrid[guardRow][guardColumn] = direction
			continue
		}

		nextRow, nextColumn := day6.FindGuardNextPosition(guardRow, guardColumn, direction)
		if roomGrid[nextRow][nextColumn] == '.' {
			distinctPositionsVisited++
		}

		// Mark visited cells with X so they don't get counted again.
		roomGrid[guardRow][guardColumn] = 'X'
		// Move guard to next cell, which we know is unblocked due to previous checks
		roomGrid[nextRow][nextColumn] = direction
	}

	return fmt.Sprint(distinctPositionsVisited)
}
