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
			Input:  "aaaaaiiiixqvsm",
			Output: "5a4ixqvsm",
		},
		{
			Input:  "rrdkuuuuyyyrrrrgghc",
			Output: "rrdk4u3y4rgghc",
		},
		{
			Input:  "xhzzzccccvvsssqppc",
			Output: "xh3z4cvv3sqppc",
		},
		{
			Input:  "jbiiiulllllvvvvtttttxxxxxs",
			Output: "jb3iu5l4v5t5xs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Input, func(t *testing.T) {
			output := Compress(tt.Input)
			if output != tt.Output {
				t.Fatalf("Compress(%s) unexpected value got: '%s', want: '%s'", tt.Input, output, tt.Output)
			}
		})
	}
}
