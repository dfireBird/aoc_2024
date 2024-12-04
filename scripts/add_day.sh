#!/usr/bin/env bash

if [[ "$#" -lt 1 ]]; then
    echo "Need a parameter for day value"
    exit 2
fi

day=$1

mkdir -p "$day"

echo 'package main

import (
	_ "embed"
	"fmt"
	"log"

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
	return 0, &internal.NotImplemented{}
}

func Part2(input string, prev_result result) (result, error) {
	return 0, &internal.NotImplemented{}
}
' > "${day}/main.go"

echo 'package main

import (
	"testing"

	"github.com/dfirebird/aoc_2024/internal"
)

type TestData internal.TestData[result]

var example = ``

func Test_part1(t *testing.T) {
	tests := []TestData{
		{
			Name:  "example",
			Input: example,
			Want:  11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got, err := Part1(tt.Input); err != nil || got != tt.Want {
				t.Errorf("part1() = %v, want %v", got, tt.Want)
			}
		})
	}
}

var example2 = ``

func Test_part2(t *testing.T) {
	tests := []TestData{
		{
			Name:  "example",
			Input: example2,
			Want:  31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got, err := Part2(tt.Input, -1); err != nil || got != tt.Want {
				t.Errorf("part2() = %v, want %v", got, tt.Want)
			}
		})
	}
}
' > "${day}/main_test.go"

echo '' > "${day}/input.txt"
