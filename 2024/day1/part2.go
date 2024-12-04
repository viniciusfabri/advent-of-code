package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var arrayA, arrayB []int

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			arrayA = append(arrayA, a)
			arrayB = append(arrayB, b)
		}
	}

	similarityScore := 0
	for _, locationId := range arrayA {
		occurrances := countOccurrences(arrayB, locationId)
		similarityScore += locationId * occurrances
	}

	return fmt.Sprint(similarityScore)
}

func countOccurrences(array []int, target int) int {
	count := 0
	for _, num := range array {
		if num == target {
			count++
		}
	}
	return count
}
