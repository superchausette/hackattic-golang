package main

import (
	"strconv"
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct {
		Input  uint
		Output string
	}{
		{
			Input:  100,
			Output: "Saturday",
		},
		{
			Input:  0,
			Output: "Thursday",
		},
		{
			Input:  128,
			Output: "Saturday",
		},
		{
			Input:  2544,
			Output: "Sunday",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.Input)), func(t *testing.T) {
			output := WhatDayWasIt(tt.Input)
			if output != tt.Output {
				t.Fatalf("WhatDayWasIt(%d) unexpected value got: '%s', want: '%s'", tt.Input, output, tt.Output)
			}
		})
	}
}
