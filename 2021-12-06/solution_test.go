package main

import "testing"

func Test_solution(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		days  int
		want  int
	}{
		{
			"sample 80 days",
			[]string{"3,4,3,1,2"},
			80,
			5934,
		},
		{
			"sample 18 days",
			[]string{"3,4,3,1,2"},
			18,
			26,
		},
		// {
		// 	"sample 256 days",
		// 	[]string{"3,4,3,1,2"},
		// 	256,
		// 	26984457539,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSolution(tt.input, tt.days); got != tt.want {
				t.Errorf("calculateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
