package shared

import (
	"slices"
	"strconv"
	"strings"
)

func ParseInputRules(input string) [][]int {
	input = strings.Trim(input, "\n")

	rules := make([][]int, 0)
	for _, rule := range strings.Split(input, "\n") {
		rule := strings.Split(rule, "|")
		pageBefore, _ := strconv.Atoi(rule[0])
		pageAfter, _ := strconv.Atoi(rule[1])
		rules = append(rules, []int{pageBefore, pageAfter})
	}
	return rules
}

func ParseInputUpdates(input string) [][]int {
	input = strings.Trim(input, "\n")

	updates := make([][]int, 0)
	for _, update := range strings.Split(input, "\n") {
		update := strings.Split(update, ",")
		updatePages := make([]int, 0)
		for _, page := range update {
			page, _ := strconv.Atoi(page)
			updatePages = append(updatePages, page)
		}
		updates = append(updates, updatePages)
	}
	return updates
}

/**
 * Find all rules that match the given pages
 */
func FindMatchingRules(rules [][]int, update []int) [][]int {
	matchingRules := make([][]int, 0)
	for _, rule := range rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
			matchingRules = append(matchingRules, rule)
		}
	}
	return matchingRules
}

func IsUpdateValid(rules [][]int, update []int) bool {
	for pageIndexInUpdate, currentPageNumber := range update {
		pagesBefore := update[:pageIndexInUpdate]
		if len(pagesBefore) > 0 {
			for _, rule := range rules {
				if rule[1] == currentPageNumber {
					// If matching page is in update, but is not BEFORE currentPage, it's invalid
					if slices.Contains(update, rule[0]) && !slices.Contains(pagesBefore, rule[0]) {
						return false
					}
				}
			}
		}

		pagesAfter := update[pageIndexInUpdate+1:]
		if len(pagesAfter) > 0 {
			for _, rule := range rules {
				if rule[0] == currentPageNumber {
					// If matching page is in update, but is not AFTER currentPage, it's invalid
					if slices.Contains(update, rule[0]) && !slices.Contains(pagesAfter, rule[1]) {
						return false
					}
				}
			}
		}
	}

	return true
}
