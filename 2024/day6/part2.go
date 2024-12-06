package main

import (
	"fmt"
	"sync"

	day6 "github.com/viniciusfabri/aoc-2024/2024/day6/shared"
)

func Part2(input string) string {
	roomGrid := day6.ParseInputIntoGrid(input)
	startRow, startCol, _ := day6.FindSecurityGuardPositionAndDirection(roomGrid)

	navigatedRoomGrid := navigateEntireGrid(day6.CopyGrid(roomGrid))
	// Find all places where an obstruction can be placed. We should only try on places the guard can actually walk on.
	// We can do this because we can only place at most one obstruction, so they will always be somewhere along the guard's original path.
	validObstructionPositions := []struct{ row, col int }{}
	for r, row := range navigatedRoomGrid {
		for c, cell := range row {
			if cell == 'X' && !(r == startRow && c == startCol) {
				validObstructionPositions = append(validObstructionPositions, struct{ row, col int }{r, c})
			}
		}
	}

	// Now for each valid obstruction position, we simulate the guard's path and check if they get stuck in a loop
	loopCount := 0
	for _, obstruction := range validObstructionPositions {
		simulatedGrid := day6.CopyGrid(roomGrid)
		simulatedGrid[obstruction.row][obstruction.col] = '#'

		if guardGetsStuckInLoop(simulatedGrid) {
			loopCount++
		}
	}

	// Alternatively you can parallelize the simulation...
	// loopCount := countLoopsConcurrently(roomGrid, validObstructionPositions, 8)

	return fmt.Sprint(loopCount)
}

/**
 * Chat GPT generated code.
 */
func countLoopsConcurrently(roomGrid [][]rune, validObstructionPositions []struct{ row, col int }, numWorkers int) string {
	type Position struct {
		row, col int
	}

	// Worker function to process obstructions
	worker := func(obstructionCh <-chan Position, resultCh chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for obstruction := range obstructionCh {
			simulatedGrid := day6.CopyGrid(roomGrid)
			simulatedGrid[obstruction.row][obstruction.col] = '#'
			if guardGetsStuckInLoop(simulatedGrid) {
				resultCh <- 1
			} else {
				resultCh <- 0
			}
		}
	}

	obstructionCh := make(chan Position, len(validObstructionPositions))
	resultCh := make(chan int, len(validObstructionPositions))

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(obstructionCh, resultCh, &wg)
	}

	// Send all obstruction positions to the channel
	go func() {
		for _, obstruction := range validObstructionPositions {
			obstructionCh <- obstruction
		}
		close(obstructionCh)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect results
	loopCount := 0
	for result := range resultCh {
		loopCount += result
	}

	return fmt.Sprint(loopCount)
}

/**
 * Takes the guard across the whole room, eventually exiting
 */
func navigateEntireGrid(roomGrid [][]rune) [][]rune {
	for {
		guardRow, guardColumn, direction := day6.FindSecurityGuardPositionAndDirection(roomGrid)

		if day6.GuardCanLeave(roomGrid, guardRow, guardColumn, direction) {
			roomGrid[guardRow][guardColumn] = 'X'
			break
		}

		if day6.HasObstaculeAhead(roomGrid, guardRow, guardColumn, direction) {
			direction = day6.NextDirection(direction)
			// Update room grid with guard facing new direction and restart loop
			roomGrid[guardRow][guardColumn] = direction
			continue
		}

		nextRow, nextColumn := day6.FindGuardNextPosition(guardRow, guardColumn, direction)
		roomGrid[guardRow][guardColumn] = 'X'
		roomGrid[nextRow][nextColumn] = direction
	}

	return roomGrid
}

func guardGetsStuckInLoop(roomGrid [][]rune) bool {
	visited := make(map[string]bool)
	for {
		guardRow, guardCol, direction := day6.FindSecurityGuardPositionAndDirection(roomGrid)

		if day6.GuardCanLeave(roomGrid, guardRow, guardCol, direction) {
			return false
		}

		// State is a combination of row + column + direction (char representing guard')
		// If we pass the same position with the same direction, we can assume we're in a loop.
		state := fmt.Sprintf("%d,%d,%c", guardRow, guardCol, direction)
		if visited[state] {
			return true
		}
		visited[state] = true

		if day6.HasObstaculeAhead(roomGrid, guardRow, guardCol, direction) {
			direction = day6.NextDirection(direction)
			// Update room grid with guard facing new direction and restart loop
			roomGrid[guardRow][guardCol] = direction
			continue
		}

		nextRow, nextCol := day6.FindGuardNextPosition(guardRow, guardCol, direction)
		roomGrid[guardRow][guardCol] = 'X' // Mark visited
		roomGrid[nextRow][nextCol] = direction
	}
}
