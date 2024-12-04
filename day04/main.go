package main

import (
	_ "embed"
	"fmt"
	"log"
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
	charSq := parseInput(input)

	count := 0

	sqLen := len(charSq)
	for x := range sqLen {
		for y := range sqLen {
			if charSq[y][x] == 'X' {
				c := countForXmasOnAllDir(charSq, x, y, sqLen)
				count += c
			}
		}
	}
	return result(count), nil
}

func Part2(input string, prev_result result) (result, error) {
	charSq := parseInput(input)

	count := 0

	sqLen := len(charSq)
	for x := range sqLen {
		for y := range sqLen {
			if charSq[y][x] == 'A' {
				c := crossPatternCheck(charSq, x, y, sqLen)
				count += c
			}
		}
	}
	return result(count), nil
}

func crossPatternCheck(charSq [][]rune, x, y, sqLen int) int {
	withInBounds := x+1 < sqLen && x-1 >= 0 && y-1 >= 0 && y+1 < sqLen

	count := 0
	if withInBounds && checkForMorS(at(charSq, x-1, y-1), at(charSq, x+1, y+1)) && checkForMorS(at(charSq, x-1, y+1), at(charSq, x+1, y-1)) {
		count += 1
	}
	return count
}

func countForXmasOnAllDir(charSq [][]rune, x, y, sqLen int) int {
	canGoRight := (x+1 < sqLen && x+2 < sqLen && x+3 < sqLen)
	canGoLeft := (x-1 >= 0 && x-2 >= 0 && x-3 >= 0)
	canGoUp := (y-1 >= 0 && y-2 >= 0 && y-3 >= 0)
	canGoDown := (y+1 < sqLen && y+2 < sqLen && y+3 < sqLen)

	count := 0
	// check right
	if canGoRight && checkMAS(at(charSq, x+1, y), at(charSq, x+2, y), at(charSq, x+3, y)) {
		count += 1
	}
	// check left
	if canGoLeft && checkMAS(at(charSq, x-1, y), at(charSq, x-2, y), at(charSq, x-3, y)) {
		count += 1
	}
	// check down
	if canGoDown && checkMAS(at(charSq, x, y+1), at(charSq, x, y+2), at(charSq, x, y+3)) {
		count += 1
	}
	// check Up
	if canGoUp && checkMAS(at(charSq, x, y-1), at(charSq, x, y-2), at(charSq, x, y-3)) {
		count += 1
	}
	// check leading
	if canGoRight && canGoDown && checkMAS(at(charSq, x+1, y+1), at(charSq, x+2, y+2), at(charSq, x+3, y+3)) {
		count += 1
	}

	// check leading rev
	if canGoLeft && canGoUp && checkMAS(at(charSq, x-1, y-1), at(charSq, x-2, y-2), at(charSq, x-3, y-3)) {
		count += 1
	}

	// check anti
	if canGoLeft && canGoDown && checkMAS(at(charSq, x-1, y+1), at(charSq, x-2, y+2), at(charSq, x-3, y+3)) {
		count += 1
	}
	// check anti rev
	if canGoRight && canGoUp && checkMAS(at(charSq, x+1, y-1), at(charSq, x+2, y-2), at(charSq, x+3, y-3)) {
		count += 1
	}

	return count
}

func at(sq [][]rune, x, y int) rune {
	return sq[y][x]
}

func checkForMorS(a, b rune) bool {
	return a == 'M' && b == 'S' || a == 'S' && b == 'M'
}

func checkMAS(a, b, c rune) bool {
	return (a == 'M' && b == 'A' && c == 'S')
}

func parseInput(input string) [][]rune {
	var charSq [][]rune

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var row []rune
		for _, rune := range line {
			row = append(row, rune)
		}

		charSq = append(charSq, row)
	}

	return charSq
}
