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
			[]string {
				"16,1,2,0,4,2,7,1,2,14",
			}, 
			37,
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
