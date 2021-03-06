package main

import "testing"

func Test_solution(t *testing.T) {

	tests := []struct {
		name     string
		input    []string
		diagonal bool
		want     int
	}{
		{
			"no input",
			[]string{},
			false,
			0,
		},
		{
			"sample",
			[]string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			false,
			5,
		},
		{
			"sample part2",
			[]string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			true,
			12,
		},
		{
			"one diagonal intersect",
			[]string{
				"8,0 -> 0,8",
				"0,0 -> 9,9",
			},
			true,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSolution(tt.input, tt.diagonal); got != tt.want {
				t.Errorf("calculateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
