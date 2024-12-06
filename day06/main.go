package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/dfirebird/aoc_2024/internal/set"
)

//go:embed input.txt
var input string

type result int

type Pos struct {
	x int
	y int
}

type Direction Pos

var UP Direction = Direction{x: 0, y: -1}
var DOWN Direction = Direction{x: 0, y: 1}
var RIGHT Direction = Direction{x: 1, y: 0}
var LEFT Direction = Direction{x: -1, y: 0}

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
	guardPos, sqDim, obstacles := parseInput(input)
	seenPos := make(set.Set[Pos])

	curPos := guardPos
	curDir := UP
	for {
		if !withinBounds(curPos, sqDim) {
			break
		}
		seenPos.Add(curPos)

		curPos, curDir = move(curPos, curDir, obstacles)
	}
	return result(len(seenPos)), nil
}

func Part2(input string, prev_result result) (result, error) {
	guardPos, sqDim, obstacles := parseInput(input)

	sum := 0
	for x := range sqDim {
		for y := range sqDim {
			curPos := Pos{x, y}
			if obstacles.Contains(curPos) || curPos == guardPos {
				continue
			}

			obstacles.Add(curPos)
			if isLooping(guardPos, sqDim, obstacles) {
				sum += 1
			}
			obstacles.Remove(curPos)
		}
	}
	return result(sum), nil
}

func isLooping(initialGuardPos Pos, sqDim int, obstacles set.Set[Pos]) bool {
	slowPos := initialGuardPos
	fastPos := initialGuardPos
	slowDir := UP
	fastDir := UP

	for {
		if !withinBounds(slowPos, sqDim) || !withinBounds(fastPos, sqDim) {
			return false
		}

		slowPos, slowDir = move(slowPos, slowDir, obstacles)
		newFastPos, newFastDir := move(fastPos, fastDir, obstacles)
		fastPos, fastDir = move(newFastPos, newFastDir, obstacles)

		if slowPos == fastPos {
			return true
		}
	}
}

func move(curPos Pos, curDir Direction, obstacles set.Set[Pos]) (Pos, Direction) {
	newPos := curPos.add(curDir)
	if obstacles.Contains(newPos) {
		curDir = curDir.turn90()
		newPos = curPos.add(curDir)
	}

	return newPos, curDir
}

func parseInput(input string) (Pos, int, set.Set[Pos]) {
	var guardPos Pos
	obstacles := make(set.Set[Pos])

	sqDim := len(strings.Split(input, "\n")[0])

	modifiedInput := strings.ReplaceAll(input, "\n", "")
	for x := range sqDim {
		for y := range sqDim {
			curPos := Pos{x, y}
			curRune := at(modifiedInput, sqDim, curPos)
			if curRune == '#' {
				obstacles.Add(curPos)
			} else if curRune == '^' {
				guardPos = curPos
			}
		}
	}

	return guardPos, sqDim, obstacles
}

func (a Pos) add(dir Direction) Pos {
	return Pos{x: a.x + dir.x, y: a.y + dir.y}
}

func (a Direction) turn90() Direction {
	switch a {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}

	panic("Unexpected direction, verify dirction variable")
}

func withinBounds(pos Pos, sqDim int) bool {
	return pos.x >= 0 && pos.x < sqDim && pos.y >= 0 && pos.y < sqDim
}

func at(s string, sqDim int, pos Pos) rune {
	return rune(s[pos.x+sqDim*pos.y])
}
