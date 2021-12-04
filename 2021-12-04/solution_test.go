package main

import "testing"

func Test_findingFirstWinningBoardsScore(t *testing.T) {

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
			"winning row",
			[]string{
				"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
				"",
				"22 13 17 11  0",
				" 8  2 23  4 24",
				"21  9 14 16  7",
				" 6 10  3 18  5",
				" 1 12 20 15 19",
				"",
				" 3 15  0  2 22",
				" 9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
				"",
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
			4512,
		},
		{
			"winning column",
			[]string{
				"7,4,23,2,16,0,14,21,10,20,8,19,13,26,1",
				"",
				"22 13 17 11  0",
				" 8  2 23  4 24",
				"21  9 14 16  7",
				" 6 10  3 18  5",
				" 1 12 20 15 19",
				"",
				" 3 15  0  2 22",
				" 9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
				"",
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
			2171,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculateSolution(tt.input); got != tt.want {
				t.Errorf("calculateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findingLastWinningBoardsScore(t *testing.T) {

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
			"winning row",
			[]string{
				"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
				"",
				"22 13 17 11  0",
				" 8  2 23  4 24",
				"21  9 14 16  7",
				" 6 10  3 18  5",
				" 1 12 20 15 19",
				"",
				" 3 15  0  2 22",
				" 9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
				"",
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
			1924,
		},
		{
			"winning column",
			[]string{
				"7,4,23,2,16,0,14,21,10,20,8,19,13,26,1",
				"",
				"22 13 17 11  0",
				" 8  2 23  4 24",
				"21  9 14 16  7",
				" 6 10  3 18  5",
				" 1 12 20 15 19",
				"",
				" 3 15  0  2 22",
				" 9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
				"",
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
			2171,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := calculateSolution(tt.input); got != tt.want {
				t.Errorf("calculateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
