package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) string {
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

	sort.Ints(arrayA)
	sort.Ints(arrayB)

	totalDifference := 0
	for i := range arrayA {
		totalDifference += abs(arrayA[i] - arrayB[i])
	}

	return fmt.Sprint(totalDifference)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
