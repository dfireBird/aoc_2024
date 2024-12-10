package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/dfirebird/aoc_2024/internal/coordinate"
	"github.com/dfirebird/aoc_2024/internal/set"
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
	topoMap, dim, trailHeads := parseInput(input)

	totalScore := 0
	for _, trailHead := range trailHeads {
		seen := set.New[coordinate.Position]()
		totalScore += walk(topoMap, dim, trailHead, &seen)
	}

	return result(totalScore), nil
}

func Part2(input string, prev_result result) (result, error) {
	topoMap, dim, trailHeads := parseInput(input)

	totalScore := 0
	for _, trailHead := range trailHeads {
		totalScore += walk(topoMap, dim, trailHead, nil)
	}

	return result(totalScore), nil
}

func walk(topoMap []int, dim int, pos coordinate.Position, seen *set.Set[coordinate.Position]) int {
	currentHeight := at(topoMap, dim, pos)
	if currentHeight == 9 {
		if seen != nil {
			if seen.Contains(pos) {
				return 0
			}
			seen.Add(pos)
		}
		return 1
	}

	score := 0
	for _, direction := range coordinate.DIRECTIONS {
		newPos := pos.Move(direction)
		if withinBounds(newPos, dim) && at(topoMap, dim, newPos)-currentHeight == 1 {
			score += walk(topoMap, dim, newPos, seen)
		}
	}

	return score
}

func at(topoMap []int, dim int, pos coordinate.Position) int {
	return topoMap[pos.X+(pos.Y*dim)]
}

func withinBounds(pos coordinate.Position, dim int) bool {
	return pos.X >= 0 && pos.X < dim && pos.Y >= 0 && pos.Y < dim
}

const INPUT_LENGTH = 1764

func parseInput(input string) ([]int, int, []coordinate.Position) {
	var topoMap []int = make([]int, 0, INPUT_LENGTH)
	var trailHeads []coordinate.Position = make([]coordinate.Position, 0, 100)
	dim := 0

	for Y, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}

		for X, c := range row {
			topoMap = append(topoMap, int(c-'0'))

			if c == '0' {
				trailHeads = append(trailHeads, coordinate.Position{X: X, Y: Y})
			}
		}
		dim++
	}

	return topoMap, dim, trailHeads
}
