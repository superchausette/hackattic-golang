package main

import (
	"strconv"
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct {
		Input  int
		Output string
	}{
		{
			Input:  8,
			Output: "8",
		},
		{
			Input:  9,
			Output: "Fizz",
		},
		{
			Input:  10,
			Output: "Buzz",
		},
		{
			Input:  11,
			Output: "11",
		},
		{
			Input:  12,
			Output: "Fizz",
		},
		{
			Input:  13,
			Output: "13",
		},
		{
			Input:  14,
			Output: "14",
		},
		{
			Input:  15,
			Output: "FizzBuzz",
		},
		{
			Input:  16,
			Output: "16",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.Input), func(t *testing.T) {
			output := FizzBuzz(tt.Input)
			if output != tt.Output {
				t.Fatalf("FizzBuzz(%d) unexpected value got: '%s', want: '%s'", tt.Input, output, tt.Output)
			}
		})
	}
}
