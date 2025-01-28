package dice

import "testing"

func TestRollDamage(t *testing.T) {
	tests := []struct {
		Input    string
		Expected []int
	}{
		{"2d4", []int{2, 8}},
		{"3d6+5", []int{8, 23}},
		{"", []int{0, 0}},
		{"8d12", []int{8, 96}},
		{"1d4-1", []int{0, 3}},
	}

	for _, tt := range tests {
		damage := SumRollDice(tt.Input)
		if damage < tt.Expected[0] || damage > tt.Expected[1] {
			t.Errorf("Roll outside of range! Got=%d, expected between %d-%d",
				damage, tt.Expected[0], tt.Expected[1])
		}
	}
}
