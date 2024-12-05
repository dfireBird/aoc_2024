package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/dfirebird/aoc_2024/internal"
	"github.com/dfirebird/aoc_2024/internal/set"
)

//go:embed input.txt
var input string

type result int

type Page int
type Update []Page
type OrderingRules map[Page][]Page

func main() {
	Result1, err := Part1(input)
	if err == nil {
		fmt.Println("Part 1 Result", Result1)
		Result2, err := Part2(input, Result1)
		if err == nil {
			fmt.Println("Part 2 Result", Result2)
		}
	} else {
		log.Fatalf("Part 1 returned an error %v", err)
	}
}

func Part1(input string) (result, error) {
	orderingRules, updates := parseInput(input)

	sum := 0
	for _, update := range updates {
		if isCorrect(orderingRules, update) {
			middlePage := update[len(update)/2]
			sum += int(middlePage)
		}
	}
	return result(sum), nil
}

func Part2(input string, prev_result result) (result, error) {
	orderingRules, updates := parseInput(input)

	cmpPages := func(a, b Page) int {
		afterPages, ok := orderingRules[a]

		if !ok {
			return 0
		}

		if slices.Contains(afterPages, b) {
			return -1
		} else {
			return 1
		}
	}

	sum := 0
	for _, update := range updates {
		if !isCorrect(orderingRules, update) {
			slices.SortStableFunc(update, cmpPages)
			middlePage := update[len(update)/2]
			sum += int(middlePage)
		}
	}
	return result(sum), nil
}

func isCorrect(rules OrderingRules, update Update) bool {
	seen := set.New[Page]()

	for _, page := range update {
		seen.Add(page)
		afterPages, ok := rules[page]

		if !ok {

			continue
		}

		for _, afterPage := range afterPages {
			if seen.Contains(afterPage) {
				return false
			}
		}
	}

	return true
}

func parseInput(input string) (OrderingRules, []Update) {
	var updates []Update
	orderingRules := OrderingRules(make(map[Page][]Page))

	splitInput := strings.Split(input, "\n\n")

	for _, ruleString := range strings.Split(splitInput[0], "\n") {
		if ruleString == "" {
			continue
		}
		splitRuleString := strings.Split(ruleString, "|")
		before := Page(internal.ToInt(splitRuleString[0]))
		after := Page(internal.ToInt(splitRuleString[1]))

		pages, ok := orderingRules[before]
		if ok {
			orderingRules[before] = append(pages, after)
		} else {
			orderingRules[before] = []Page{after}
		}
	}

	for _, updateString := range strings.Split(splitInput[1], "\n") {
		if updateString == "" {
			continue
		}
		update := internal.Map(strings.Split(updateString, ","),
			func(u string) Page { return Page(internal.ToInt(u)) })

		updates = append(updates, Update(update))
	}

	return orderingRules, updates
}
