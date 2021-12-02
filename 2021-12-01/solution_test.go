package main

import "testing"

func Test_countIncreasingMeasurements(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"no input",
			[]int{},
			0,
		},
		{
			"small input",
			[]int{42},
			0,
		},
		{
			"sample input #1",
			[]int{1, 2, 2, 42},
			2,
		},
		{
			"increasing input",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			9,
		},
		{
			"decreasing input",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			0,
		},
		{
			"sample input",
			[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreasingMeasurements(tt.input); got != tt.want {
				t.Errorf("countIncreasingMeasurements() = %v, want %v", got, tt.want)
			}
		})
	}
}
