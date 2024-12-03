package main

import (
	"testing"

	"github.com/dfirebird/aoc_2024/internal"
)

type TestData internal.TestData[result]

var example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func Test_part1(t *testing.T) {
	tests := []TestData{
		{
			Name:  "example",
			Input: example,
			Want:  2,
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

var example2 = example

func Test_part2(t *testing.T) {
	tests := []TestData{
		{
			Name:  "example",
			Input: example2,
			Want:  4,
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
