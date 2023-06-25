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
			Input:  "szWindowContents",
			Output: "window_contents",
		},
		{
			Input:  "iAirflowParameter",
			Output: "airflow_parameter",
		},
		{
			Input:  "fMixtureRatio",
			Output: "mixture_ratio",
		},
		{
			Input:  "i32AirflowForce",
			Output: "airflow_force",
		},
		{
			Input:  "air2Param",
			Output: "air2_param",
		},
		{
			Input:  "i16AirDangerLevel",
			Output: "air_danger_level",
		},
		{
			Input:  "bMixture2WarnLevel",
			Output: "mixture2_warn_level",
		},
		{
			Input:  "u64MixtureParam",
			Output: "mixture_param",
		},
		{
			Input:  "u32Throttle2Mass",
			Output: "throttle2_mass",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Input, func(t *testing.T) {
			output := ToSnakeCase(tt.Input)
			if output != tt.Output {
				t.Fatalf("ToSnakeCase(%s) unexpected value got: '%s', want: '%s'", tt.Input, output, tt.Output)
			}
		})
	}
}
