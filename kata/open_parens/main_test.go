package main

import (
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "(())",
			Output: true,
		},
		{
			Input:  "()))",
			Output: false,
		},
		{
			Input:  "(()((())))",
			Output: true,
		},
		{
			Input:  "(()(()(()))",
			Output: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Input, func(t *testing.T) {
			output := AreParensProperlyNested(tt.Input)
			if output != tt.Output {
				t.Fatalf("AreParensProperlyNested(%s) unexpected value got: '%t', want: '%t'", tt.Input, output, tt.Output)
			}
		})
	}
}
