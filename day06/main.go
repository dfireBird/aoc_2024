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
	guardPos, sqDim, obstacles := parseInput(input)
	seenPos := make(set.Set[coordinate.Position])

	curPos := guardPos
	curDir := coordinate.UP
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

	seenPos := make(set.Set[coordinate.Position])

	curPos := guardPos
	curDir := coordinate.UP
	for {
		if !withinBounds(curPos, sqDim) {
			break
		}
		seenPos.Add(curPos)

		curPos, curDir = move(curPos, curDir, obstacles)
	}

	sum := 0
	for curPos := range seenPos {
		if obstacles.Contains(curPos) || curPos == guardPos {
			continue
		}

		obstacles.Add(curPos)
		if isLooping(guardPos, sqDim, obstacles) {
			sum += 1
		}
		obstacles.Remove(curPos)

	}

	return result(sum), nil
}

func isLooping(initialGuardPos coordinate.Position, sqDim int, obstacles set.Set[coordinate.Position]) bool {
	type PosDir struct {
		pos coordinate.Position
		dir coordinate.Direction
	}

	seen := make(set.Set[PosDir])

	curPosDir := PosDir{initialGuardPos, coordinate.UP}
	for {
		if !withinBounds(curPosDir.pos, sqDim) {
			return false
		}

		if seen.Contains(curPosDir) {
			return true
		}

		seen.Add(curPosDir)

		curPosDir.pos, curPosDir.dir = move(curPosDir.pos, curPosDir.dir, obstacles)
	}
}

func move(curPos coordinate.Position, curDir coordinate.Direction, obstacles set.Set[coordinate.Position]) (coordinate.Position, coordinate.Direction) {
	newPos := curPos.Move(curDir)
	if obstacles.Contains(newPos) {
		return curPos, curDir.Turn90()
	}

	return newPos, curDir
}

func parseInput(input string) (coordinate.Position, int, set.Set[coordinate.Position]) {
	var guardPos coordinate.Position
	obstacles := make(set.Set[coordinate.Position])

	sqDim := len(strings.Split(input, "\n")[0])

	modifiedInput := strings.ReplaceAll(input, "\n", "")
	for x := range sqDim {
		for y := range sqDim {
			curPos := coordinate.Position{x, y}
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

func withinBounds(pos coordinate.Position, sqDim int) bool {
	return pos.X >= 0 && pos.X < sqDim && pos.Y >= 0 && pos.Y < sqDim
}

func at(s string, sqDim int, pos coordinate.Position) rune {
	return rune(s[pos.X+sqDim*pos.Y])
}
