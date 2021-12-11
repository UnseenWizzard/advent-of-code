package main

import "testing"

func Test_part2(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			"sample",
			[]string{
				"5877777777",
				"8877777777",
				"7777777777",
				"7777777777",
				"7777777777",
				"7777777777",
				"7777777777",
				"7777777777",
				"7777777777",
				"7777777777",
			},
			2,
		},
		{
			"sample",
			[]string{
				"47666",
				"77666",
				"66666",
				"66666",
				"66666",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePartTwo(tt.input); got != tt.want {
				t.Errorf("calculatePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
