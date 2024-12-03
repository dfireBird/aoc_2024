package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type result int

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
	reports := parseInput(input)

	noOfSafe := 0

	for _, report := range reports {
		if isSafe(report) {
			noOfSafe += 1
		}
	}

	return result(noOfSafe), nil
}

func Part2(input string, prev_result result) (result, error) {
	reports := parseInput(input)

	noOfSafe := 0

	for _, report := range reports {
		if isSafe(report) {
			noOfSafe += 1
			continue
		}

		for i := range report {
			removed := copyAndRemove(report, i)
			if isSafe(removed) {
				noOfSafe += 1
				break
			}
		}
	}

	return result(noOfSafe), nil
}

func parseInput(input string) [][]int {
	var reports [][]int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var report []int
		for _, v := range strings.Split(line, " ") {
			level, _ := strconv.Atoi(v)

			report = append(report, level)
		}

		reports = append(reports, report)
	}
	return reports
}

func isSafe(report []int) bool {
	diff := report[1] - report[0]

	if abs(diff) > 3 || abs(diff) < 1 {
		return false
	}

	isSafe := true
	for i := 2; i < len(report); i++ {
		curDiff := report[i] - report[i-1]

		if abs(curDiff) > 3 || abs(curDiff) < 1 || (curDiff > 0) != (diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}

func copyAndRemove(slice []int, idx int) []int {
	newSlice := make([]int, len(slice)-1)
	newSliceIdx := 0
	for i, val := range slice {
		if i == idx {
			continue
		}

		newSlice[newSliceIdx] = val
		newSliceIdx++
	}

	return newSlice
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
