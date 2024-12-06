package day6

import "strings"

func ParseInputIntoGrid(input string) [][]rune {
	rows := strings.Split(input, "\n")
	roomGrid := make([][]rune, len(rows))
	for i, row := range rows {
		roomGrid[i] = []rune(row)
	}

	return roomGrid
}

func FindSecurityGuardPositionAndDirection(roomGrid [][]rune) (int, int, rune) {
	for i, row := range roomGrid {
		for j, cell := range row {
			if cell == '^' || cell == '>' || cell == 'V' || cell == '<' {
				return i, j, cell
			}
		}
	}

	return -1, -1, 'x'
}

func GuardCanLeave(roomGrid [][]rune, row, column int, direction rune) bool {
	// If guard is looking up, check if they are on the top row
	if direction == '^' && row == 0 {
		return true
	}

	// If guard is looking right, check if they are on the rightmost column
	if direction == '>' && column == len(roomGrid[0])-1 {
		return true
	}

	// If guard is looking down, check if they are on the bottom row
	if direction == 'V' && row == len(roomGrid)-1 {
		return true
	}

	// If guard is looking left, check if they are on the leftmost column
	if direction == '<' && column == 0 {
		return true
	}

	return false
}

func HasObstaculeAhead(roomGrid [][]rune, row, column int, direction rune) bool {
	if direction == '^' && row > 0 && roomGrid[row-1][column] == '#' {
		return true
	}

	if direction == '>' && column < len(roomGrid[0])-1 && roomGrid[row][column+1] == '#' {
		return true
	}

	if direction == 'V' && row < len(roomGrid)-1 && roomGrid[row+1][column] == '#' {
		return true
	}

	if direction == '<' && column > 0 && roomGrid[row][column-1] == '#' {
		return true
	}

	return false
}

func NextDirection(currentDirection rune) rune {
	switch currentDirection {
	case '^':
		return '>'
	case '>':
		return 'V'
	case 'V':
		return '<'
	case '<':
		return '^'
	}

	return 'x'
}

func FindGuardNextPosition(currentRow, currentColumn int, currentDirection rune) (int, int) {
	switch currentDirection {
	case '^':
		return currentRow - 1, currentColumn
	case '>':
		return currentRow, currentColumn + 1
	case 'V':
		return currentRow + 1, currentColumn
	case '<':
		return currentRow, currentColumn - 1
	}

	return -1, -1
}

func CopyGrid(grid [][]rune) [][]rune {
	copied := make([][]rune, len(grid))
	for i := range grid {
		copied[i] = make([]rune, len(grid[i]))
		copy(copied[i], grid[i])
	}
	return copied
}
