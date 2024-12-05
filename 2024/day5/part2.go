package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/viniciusfabri/aoc-2024/2024/day5/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	blocks := strings.Split(input, "\n\n")
	rules := shared.ParseInputRules(blocks[0])
	updates := shared.ParseInputUpdates(blocks[1])

	invalidUpdates := make([][]int, 0)
	for _, update := range updates {
		matchingRules := shared.FindMatchingRules(rules, update)
		isUpdateValid := shared.IsUpdateValid(matchingRules, update)

		if !isUpdateValid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	newUpdates := make([][]int, 0)
	for _, invalidUpdate := range invalidUpdates {
		matchingRules := shared.FindMatchingRules(rules, invalidUpdate)
		fixedUpdate := fixInvalidUpdate(matchingRules, invalidUpdate)

		newUpdates = append(newUpdates, fixedUpdate)
	}

	sumOfMiddlePages := 0
	for _, validUpdate := range newUpdates {
		updateLength := len(validUpdate)
		sumOfMiddlePages += validUpdate[updateLength/2]
	}

	return fmt.Sprint(sumOfMiddlePages)
}

func fixInvalidUpdate(rules [][]int, invalidUpdate []int) []int {
	newUpdate := invalidUpdate
	updated := true
	for updated {
		updated = false
		for i := 0; i < len(invalidUpdate); i++ {
			for _, rule := range rules {
				posA := slices.Index(invalidUpdate, rule[0])
				posB := slices.Index(invalidUpdate, rule[1])

				if posA > posB {
					// Swap characters
					newUpdate[posA], newUpdate[posB] = invalidUpdate[posB], invalidUpdate[posA]
					// Set updated to true so we run through it again
					updated = true
				}
			}
		}
	}

	return newUpdate
}
