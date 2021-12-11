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
			"valid 1",
			[]string{"()"},
			0,
		},
		{
			"valid 2",
			[]string{"{()()()}"},
			0,
		},
		{
			"valid 3",
			[]string{"[<>({}){}[([])<>]]"},
			0,
		},
		{
			"valid 4",
			[]string{"(((((((((())))))))))"},
			0,
		},
		{
			"invalid 1",
			[]string{"(]"},
			57,
		},
		{
			"invalid 2",
			[]string{"{()()()>"},
			25137,
		},
		{
			"invalid 3",
			[]string{"<([]){()}[{}])"},
			3,
		},
		{
			"invalid 4",
			[]string{"<([]){()}[{}}(>"},
			1197,
		},
		{
			"invalid 4",
			[]string{
				"<([]){()}[{}])", 
				"(]>",
				">", 
			},
			3 + 57 + 25137,
		},
		{
			"sample",
			[]string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"{([(<{}[<>[]}>{[]{[(<()>",
				"(((({<>}<{<{<>}{[]{[]{}",
				"[[<[([]))<([[{}[[()]]]",
				"[{[{({}]{}}([{[{{{}}([]",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"[<(<(<(<{}))><([]([]()",
				"<{([([[(<>()){}]>(<<{{",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
			26397,
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
