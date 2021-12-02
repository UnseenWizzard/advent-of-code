package main

import "testing"

func Test_countIncreasingMeasurements(t *testing.T) {

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
			"sample input",
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSolution(applyCommands(tt.input)); got != tt.want {
				t.Errorf("countIncreasingMeasurements() = %v, want %v", got, tt.want)
			}
		})
	}
}
