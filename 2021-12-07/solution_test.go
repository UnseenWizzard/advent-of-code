package main

import "testing"

func Test_solution(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		complexFuelCost bool
		want  int
	}{
		{
			"no input",
			[]string{},
			false,
			0,
		
		},
		{
			"sample",
			[]string {
				"16,1,2,0,4,2,7,1,2,14",
			}, 
			false,
			37,
		},
		{
			"part 2 sample",
			[]string {
				"16,1,2,0,4,2,7,1,2,14",
			}, 
			true,
			168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSolution(tt.input, tt.complexFuelCost); got != tt.want {
				t.Errorf("calculateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
