package main

import (
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct {
		Input  string
		Output uint16
	}{
		{
			Input:  "#.#.#.###.#.##.#",
			Output: 43949,
		},
		{
			Input:  "##.##.......#..#",
			Output: 55305,
		},
		{
			Input:  "#..#####..#.#...",
			Output: 40744,
		},
		{
			Input:  "###..#....###.##",
			Output: 58427,
		},
		{
			Input:  "#..#..#.#....##",
			Output: 18755,
		},
		{
			Input:  "#......#.##....",
			Output: 16560,
		},
		{
			Input:  "#############.#.",
			Output: 65530,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Input, func(t *testing.T) {
			output := DecodeAlmostBinary(tt.Input)
			if output != tt.Output {
				t.Fatalf("DecodeAlmostBinary(%s) unexpected value got: '%d', want: '%d'", tt.Input, output, tt.Output)
			}
		})
	}
}
