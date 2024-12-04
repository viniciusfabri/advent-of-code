package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Exported function to be called by the main application
func Part1(input string) string {
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

		isBigger := true
		isSmaller := true
		withinRange := true
		for i := 1; i < len(values); i++ {
			previous := values[i-1]
			current := values[i]

			if isSmaller && previous > current {
				isSmaller = false
			}

			if isBigger && previous < current {
				isBigger = false
			}

			difference := abs(previous - current)
			// within range = Any two adjacent levels differ by at least one and at most three
			if withinRange && (difference < 1 || difference > 3) {
				withinRange = false
			}
		}
		// If values are still increasing or decreasing, and within range, then the report is safe
		if (isBigger && withinRange) || (isSmaller && withinRange) {
			safeReports++
		}
	}

	return fmt.Sprint(safeReports)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
