package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
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
	leftList, rightList := parseInput(input)
	sumDistance := 0

	slices.Sort(leftList)
	slices.Sort(rightList)

	for i := range leftList {
		sumDistance += distance(leftList[i], rightList[i])
	}

	return result(sumDistance), nil
}

func Part2(input string, prev_result result) (result, error) {
	leftList, rightList := parseInput(input)

	frequencyMap := calculateFreq(rightList)
	similarityScore := 0

	for _, v := range leftList {
		frequency := frequencyMap[v]

		similarityScore += v * frequency // zero value if not present
	}

	return result(similarityScore), nil
}

func parseInput(input string) ([]int, []int) {
	var leftList, rightList []int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		splitV := strings.Split(line, "   ")

		// Will not Error
		leftElem, _ := strconv.Atoi(splitV[0])
		rightElem, _ := strconv.Atoi(splitV[1])

		leftList = append(leftList, leftElem)
		rightList = append(rightList, rightElem)
	}

	return leftList, rightList
}

func calculateFreq(slice []int) map[int]int {
	frequencyMap := make(map[int]int)
	for _, v := range slice {
		elem, ok := frequencyMap[v]
		if !ok {
			frequencyMap[v] = 1
		} else {
			frequencyMap[v] = elem + 1
		}
	}

	return frequencyMap
}

func distance(a, b int) int {
	if diff := a - b; diff < 0 {
		return -diff
	} else {
		return diff
	}
}
