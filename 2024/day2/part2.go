package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Exported function to be called by the main application
func Part2(input string) string {
	report := strings.Split(strings.TrimSpace(input), "\n")

	safeReports := 0
	for _, line := range report {
		parts := strings.Fields(line)
		var values []int
		for _, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				fmt.Print(err)
				continue
			}
			values = append(values, value)
		}

		if isSafe(values, false) || isSafe(values, true) {
			safeReports++

		}
	}

	return fmt.Sprint(safeReports)
}

func isSafe(values []int, tolerate bool) bool {
	var sortAsc bool

	for i := 0; i < len(values); i++ {
		faulty := false

		levels := values

		if tolerate {
			levels = make([]int, 0)
			levels = append(levels, values[:i]...)
			levels = append(levels, values[i+1:]...)
		}

		// Determine sorting
		if levels[0] < levels[1] {
			sortAsc = true
		} else {
			sortAsc = false
		}

		for i := 1; i < len(levels); i++ {
			if sortAsc {
				if levels[i] < levels[i-1] {
					faulty = true
				}
			} else {
				if levels[i] > levels[i-1] {
					faulty = true
				}
			}

			differ := ab2(levels[i] - levels[i-1])
			if differ < 1 || differ > 3 {
				faulty = true
			}
		}

		if !faulty {
			return true
		}

		if !tolerate {
			return false
		}
	}

	return false
}

func ab2(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
