package utils

import (
	"slices"
	"testing"
)

func TestSanitizeArgs(t *testing.T) {
	type Test struct {
		Input    []string
		Expected []string
	}

	tests := []Test{
		{
			Input:    []string{" Grob ", " -a"},
			Expected: []string{"Grob", "-a"},
		},
		{
			Input:    []string{" -Ab", "   Funky Frank         "},
			Expected: []string{"-Ab", "Funky Frank"},
		},
		{
			Input:    []string{"      -i     ", " ", " "},
			Expected: []string{"-i", "", ""},
		},
	}

	for _, tt := range tests {
		args := SanitizeArgs(tt.Input)
		if !slices.Equal(args, tt.Expected) {
			t.Errorf("Oops! Got %v, expected %v", args, tt.Expected)
		}
	}
}
