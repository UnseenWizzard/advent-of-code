package main

import (
	"log"
	"os"
	"strings"

	"riedmann.dev/aoc21/util"
)

func parseInput(input []string) (observations [][]string, outputs [][]string) {
	observations = make([][]string, len(input))
	outputs = make([][]string, len(input))
	for i, l := range input {
		s := strings.Split(l, "|")
		observations[i] = strings.Fields(s[0])
		outputs[i] = strings.Fields(s[1])	
	}
	return
}

func countKnownDigits(displays []string) (count int) {
	for _, d := range displays {
		activeSegments := len(d)
		if activeSegments != 5 && activeSegments != 6 {
			//1, 4, 7 or 8 (2,4,3,7 active segments)
			count++
		} 
	}
	return
}

func calculateSolution(input []string) int {
    _, outputs := parseInput(input)
	count := 0
	for _, o := range outputs {
		count += countKnownDigits(o)
	}
    return count
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Session Cookie needed as arg to get puzzle input!")
    }
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input))
}
const day = "8"
