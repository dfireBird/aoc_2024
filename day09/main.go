package main

import (
	_ "embed"
	"fmt"
	"log"
	"maps"
	"slices"

	"github.com/dfirebird/aoc_2024/internal"
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
	fsMap, _, _, _ := parseInput(input)
	fsMap = defrag(fsMap)
	return result(checksum(fsMap)), nil
}

func Part2(input string, prev_result result) (result, error) {
	fsMap, _, fileMap, freeMap := parseInput(input)
	fsMap = compact(fsMap, fileMap, freeMap)
	return result(checksum(fsMap)), nil
}

func checksum(fs []int) int {
	sum := 0
	for i, val := range fs {
		if val == -1 {
			continue
		}

		sum += i * val
	}

	return sum
}

func defrag(fs []int) []int {
	copiedFs := append(make([]int, 0, len(fs)), fs...)

	var fileBlockIdxs []int
	for i, val := range fs {
		if val != -1 {
			fileBlockIdxs = append(fileBlockIdxs, i)
		}
	}

	idx := len(fileBlockIdxs) - 1
	for i, block := range copiedFs {
		if idx < 0 || isDefragDone(fs) {
			break
		}

		if block == -1 {
			fs[i], fs[fileBlockIdxs[idx]] = fs[fileBlockIdxs[idx]], fs[i]
			idx--
		}
	}
	return fs
}

func isDefragDone(fs []int) bool {
	var freeIdxs []int

	for i, val := range fs {
		if val == -1 {
			freeIdxs = append(freeIdxs, i)
		}
	}

	var prevIdx = freeIdxs[0]
	for i, idx := range freeIdxs {
		if i == 0 {
			continue
		}

		if idx-prevIdx > 1 {
			return false
		}

		prevIdx = idx
	}

	return true
}

func compact(fs []int, fileMap, freeMap map[int]int) []int {
	sortedFileKeys := slices.Sorted(maps.Keys(fileMap))
	sortedFreeKeys := slices.Sorted(maps.Keys(freeMap))

	slices.Reverse(sortedFileKeys)

	lastIdx := len(sortedFileKeys) - 1

	for i, startIdx := range sortedFileKeys {
		fileIdx := lastIdx - i
		length := fileMap[startIdx]

		idxForKey := slices.IndexFunc(sortedFreeKeys, func(e int) bool {
			return e < startIdx && freeMap[e] >= length
		})
		if idxForKey == -1 {
			fileIdx--
			continue
		}

		neededFreeIdx := sortedFreeKeys[idxForKey]

		originalSize := freeMap[neededFreeIdx]
		delete(freeMap, neededFreeIdx)

		copiedFileStartIdx := startIdx
		for range length {
			fs[neededFreeIdx] = fileIdx
			fs[copiedFileStartIdx] = -1
			copiedFileStartIdx++
			neededFreeIdx++
		}

		freeMap[neededFreeIdx] = originalSize - length
		sortedFreeKeys = slices.Sorted(maps.Keys(freeMap))
		fileIdx--
	}

	return fs
}

func printFS(fs []int) {
	for _, val := range fs {
		if val == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(val)
		}
	}
	fmt.Print("\n")
}

func parseInput(input string) ([]int, int, map[int]int, map[int]int) {
	fsLength := 0

	var fsMap []int

	fileMap := make(map[int]int)
	freeMap := make(map[int]int)

	fileIdx := 0
	for i, value := range input {
		if value == '\n' {
			continue
		}

		valueInInt := internal.ToInt(string(value))
		startIdx := len(fsMap)
		for range valueInInt {
			var appendVal int
			if i%2 == 0 {
				appendVal = fileIdx
			} else {
				appendVal = -1
			}
			fsMap = append(fsMap, appendVal)
		}

		if i%2 == 0 {
			fileIdx++
			fileMap[startIdx] = valueInInt
		} else {
			freeMap[startIdx] = valueInInt
		}

		fsLength += valueInInt
	}

	return fsMap, fsLength, fileMap, freeMap
}
