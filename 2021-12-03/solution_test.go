package main

import "testing"

func Test_solution(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			"no input",
			[]string{},
			0,
		},
		{
			"sample",
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSolution(tt.input); got != tt.want {
				t.Errorf("calculateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
