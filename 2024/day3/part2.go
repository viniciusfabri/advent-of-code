package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Exported function to be called by the main application
func Part2(input string) string {
	// Regular expressions for matching patterns
	mulRegex := regexp.MustCompile(`^mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`^do\(\)`)
	dontRegex := regexp.MustCompile(`^don't\(\)`)

	// Variable to store the sum of all multiplications
	sum := 0

	// Initially, mul instructions are enabled
	mulEnabled := true

	// Iterate through the input, treating it as a stream of characters
	for i := 0; i < len(input); {
		// Match `do()` or `don't()`
		if doRegex.MatchString(input[i:]) {
			mulEnabled = true
			i += len("do()") // Skip the length of the match
		} else if dontRegex.MatchString(input[i:]) {
			mulEnabled = false
			i += len("don't()") // Skip the length of the match
		} else {
			// Match `mul(x, y)`
			mulMatch := mulRegex.FindStringSubmatchIndex(input[i:])
			if mulMatch != nil && mulEnabled {
				// Extract numbers and compute the multiplication
				start, end := mulMatch[0], mulMatch[1]
				numbers := mulRegex.FindStringSubmatch(input[i+start : i+end])
				num1, _ := strconv.Atoi(numbers[1])
				num2, _ := strconv.Atoi(numbers[2])
				sum += num1 * num2
				i += end // Move the pointer past the match
			} else {
				// Move to the next character if no match
				i++
			}
		}
	}

	return fmt.Sprint(sum)
}
