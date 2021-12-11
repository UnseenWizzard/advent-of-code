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
			"no input",
			[]string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			15,
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
