package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dfirebird/aoc_2024/internal"
)

//go:embed input.txt
var input string

type result int

type Eq struct {
	ans      int
	operands []int
}

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
	equations := parseInput(input)
	sum := 0
	for _, equation := range equations {
		if equation.isPossible(false) {
			sum += equation.ans
		}
	}
	return result(sum), nil
}

func Part2(input string, prev_result result) (result, error) {
	equations := parseInput(input)
	sum := 0
	for _, equation := range equations {
		if equation.isPossible(true) {
			sum += equation.ans
		}
	}
	return result(sum), nil
}

func (eq Eq) isPossible(concatenate bool) bool {
	return eq.isPossibleDriver(1, eq.operands[0], concatenate)
}

func (eq Eq) isPossibleDriver(idx, accumulate int, concatenate bool) bool {
	if eq.ans == accumulate {
		return true
	}

	if idx >= len(eq.operands) || accumulate > eq.ans {
		return false
	}

	concatenateResult := false
	if concatenate {
		concatenateResult = eq.isPossibleDriver(idx+1, digitConcatenate(accumulate, eq.operands[idx]), concatenate)
	}

	return eq.isPossibleDriver(idx+1, accumulate+eq.operands[idx], concatenate) || eq.isPossibleDriver(idx+1, accumulate*eq.operands[idx], concatenate) || concatenateResult
}

func parseInput(input string) []Eq {
	var equations []Eq
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		splitLine := strings.Split(line, ": ")
		ans := internal.ToInt(splitLine[0])
		operands := internal.Map(strings.Split(splitLine[1], " "), internal.ToInt)

		equations = append(equations, Eq{ans, operands})
	}

	return equations
}

func digitConcatenate(a, b int) int {
	return internal.ToInt(strconv.Itoa(a) + strconv.Itoa(b))
}
