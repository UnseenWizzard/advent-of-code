package main

import (
	"os"
	"log"
	"riedmann.dev/aoc21/util"
)

func findLargest(in []int) (largest int) {
	for _, i := range in {
		if i > largest {
			largest = i
		}
	}
	return
}

func calcDistance(from int, to int) int {
	if from > to {
		return from - to
	}
	return to - from
}

func calculateSolution(input []string) int {
	if len(input) != 1 {
		println("Expected single line input, but got ", len(input))
		return 0
	}

    depths := util.ParseIntList(input[0])
	largestDepth := findLargest(depths)
	lowest := int(^uint(0) >>1) //max signed int val
	
	for i := 0; i <= largestDepth; i++ {
		cost := 0
		for _, d := range depths {
			cost += calcDistance(i, d)
			
		}
		if cost < lowest {
			lowest = cost
		}
	}
    return lowest
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Session Cookie needed as arg to get puzzle input!")
    }
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input))
}
const day = "7"
