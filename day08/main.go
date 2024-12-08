package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"unicode"

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
	sqDim, antennas := parseInput(input)

	uniqueAntiNodeLocations := set.New[coordinate.Position]()

	for _, locations := range antennas {

		for _, permutedLocation := range permutations(locations) {
			posA, posB := findAntiNodes(permutedLocation.A, permutedLocation.B)

			if withinBound(posA, sqDim) {
				uniqueAntiNodeLocations.Add(posA)
			}

			if withinBound(posB, sqDim) {
				uniqueAntiNodeLocations.Add(posB)
			}
		}
	}

	return result(len(uniqueAntiNodeLocations)), nil
}

func Part2(input string, prev_result result) (result, error) {
	sqDim, antennas := parseInput(input)

	uniqueAntiNodeLocations := set.New[coordinate.Position]()

	for _, locations := range antennas {

		for _, permutedLocation := range permutations(locations) {
			antiNodePositions := findAntiNodesPart2(permutedLocation.A, permutedLocation.B, sqDim)

			for _, pos := range antiNodePositions {
				if withinBound(pos, sqDim) {
					uniqueAntiNodeLocations.Add(pos)
				}
			}
		}
	}

	return result(len(uniqueAntiNodeLocations)), nil
}

func findAntiNodes(a, b coordinate.Position) (coordinate.Position, coordinate.Position) {
	return findPosAtDist(a, b, 2), findPosAtDist(b, a, 2)
}
func findAntiNodesPart2(a, b coordinate.Position, sqDim int) []coordinate.Position {
	return append(findAllPosAtSameLine(a, b, sqDim), findAllPosAtSameLine(b, a, sqDim)...)
}

func findAllPosAtSameLine(a, b coordinate.Position, sqDim int) []coordinate.Position {
	var result []coordinate.Position

	dist := 1
	for {
		pos := findPosAtDist(a, b, dist)
		if !withinBound(pos, sqDim) {
			return result
		}

		result = append(result, pos)
		dist += 1
	}
}

func findPosAtDist(a, b coordinate.Position, dist int) coordinate.Position {
	return coordinate.Position{
		X: ((1 - dist) * a.X) + (dist * b.X),
		Y: ((1 - dist) * a.Y) + (dist * b.Y),
	}
}

func parseInput(input string) (int, map[string]([]coordinate.Position)) {
	sqDim := 0
	antennas := make(map[string]([]coordinate.Position))

	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for x, c := range line {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				v, ok := antennas[string(c)]
				if !ok {
					v = []coordinate.Position{}
				}

				antennas[string(c)] = append(v, coordinate.Position{X: x, Y: y})
			}
		}
		sqDim += 1
	}
	return sqDim, antennas
}

func withinBound(a coordinate.Position, sqDim int) bool {
	return a.X >= 0 && a.X < sqDim && a.Y >= 0 && a.Y < sqDim
}

func permutations[T any](slice []T) []struct {
	A T
	B T
} {
	var result []struct {
		A T
		B T
	}
	for i, a := range slice {
		for j, b := range slice {
			if i == j {
				continue
			}

			result = append(result, struct {
				A T
				B T
			}{a, b})
		}
	}

	return result
}
