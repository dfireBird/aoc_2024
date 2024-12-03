package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type result int

var mulOnly = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
var mulWithDoDont = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

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
	input = cleanInput(input)

	sum := 0
	for _, val := range mulOnly.FindAllStringSubmatch(input, -1) {
		a, _ := strconv.Atoi(val[1])
		b, _ := strconv.Atoi(val[2])
		sum += (a * b)
	}
	return result(sum), nil
}

func Part2(input string, prev_result result) (result, error) {
	input = cleanInput(input)
	isEnabled := true
	sum := 0
	for _, val := range mulWithDoDont.FindAllStringSubmatch(input, -1) {
		if val[0] == "do()" {
			isEnabled = true
		} else if val[0] == "don't()" {
			isEnabled = false
		} else {
			if isEnabled {
				a, _ := strconv.Atoi(val[1])
				b, _ := strconv.Atoi(val[2])
				sum += (a * b)
			}
		}
	}
	return result(sum), nil
}

func cleanInput(input string) string {
	return strings.ReplaceAll(input, "\n", "")
}
