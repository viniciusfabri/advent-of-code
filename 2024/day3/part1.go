package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Exported function to be called by the main application
func Part1(input string) string {
	// Regular expression to match the pattern "mul(number,number)"
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	// Variable to store the sum of all multiplications
	sum := 0

	// Iterate over the matches
	for _, match := range matches {
		// Extract the numbers from the match
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		// Multiply the numbers and add to the sum
		sum += num1 * num2
	}

	return fmt.Sprint(sum)
}
