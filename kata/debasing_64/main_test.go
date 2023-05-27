package main

import (
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "bGF0ZS1hdC1uaWdodA==",
			Output: "late-at-night",
		},
		{
			Input:  "d2l0aC10aGUtcmlzaW5nLWFwZQ==",
			Output: "with-the-rising-ape",
		},
		{
			Input:  "dGhlLXJ1dGhsZXNzLXNldmVu",
			Output: "the-ruthless-seven",
		},
	}
	decodeMap := CreateDecodeMap()
	for _, tt := range tests {
		t.Run(tt.Input, func(t *testing.T) {
			output := Debase64(tt.Input, decodeMap)
			if output != tt.Output {
				t.Fatalf("Debase64(%s) unexpected value got: '%s', want: '%s'", tt.Input, output, tt.Output)
			}
		})
	}
}
