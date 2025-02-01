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
			t.Errorf("Oops! Got %s, expected %s", args, tt.Expected)
		}
	}
}

func TestTruncateString(t *testing.T) {
	type Test struct {
		S        string
		Max      int
		Expected string
	}

	tests := []Test{
		{
			S:        "This man.",
			Max:      1,
			Expected: "T",
		},
		{
			S:        "Hi!",
			Max:      20,
			Expected: "Hi!",
		},
		{S: "",
			Max:      10,
			Expected: "",
		},
		{
			S:        "Banana Grabber!",
			Max:      0,
			Expected: "",
		},
	}

	for _, tt := range tests {
		result := TruncateString(tt.S, tt.Max)
		if result != tt.Expected {
			t.Errorf("Oops! got %s, expected %s", result, tt.Expected)
		}
	}
}
