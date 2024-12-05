package main

import (
	"fmt"
	"strings"

	"github.com/viniciusfabri/aoc-2024/2024/day5/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	blocks := strings.Split(input, "\n\n")
	rules := shared.ParseInputRules(blocks[0])
	updates := shared.ParseInputUpdates(blocks[1])

	validUpdates := make([][]int, 0)
	for _, update := range updates {
		matchingRules := shared.FindMatchingRules(rules, update)
		isUpdateValid := shared.IsUpdateValid(matchingRules, update)

		if isUpdateValid {
			validUpdates = append(validUpdates, update)
		}
	}

	sumOfMiddlePages := 0
	for _, validUpdate := range validUpdates {
		updateLength := len(validUpdate)
		sumOfMiddlePages += validUpdate[updateLength/2]
	}

	return fmt.Sprint(sumOfMiddlePages)
}
