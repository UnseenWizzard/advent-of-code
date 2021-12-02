package main

import (
	"os"
	"riedmann.dev/aoc21/util"
	"strconv"
)

// how many meassurements are greater than the previous one?
func countIncreasingMeasurements(measurements []int) int {
	if len(measurements) < 2 {
		return 0
	}

	increases := 0
	prev := measurements[0]
	for _, m := range measurements {
		if m > prev {
			increases++
		}
		prev = m
	}
	return increases
}

func countSlidingWindowIncreases(measurements []int) int {
	if len(measurements) < 3 {
		return 0
	}

	increases := 0
	prev := measurements[0] + measurements[1] + measurements[2]
	for i := range measurements {
		if i+2 >= len(measurements) {
			return increases
		}
		m := measurements[i] + measurements[i+1] + measurements[i+2]
		if m > prev {
			increases++
		}
		prev = m
	}
	return increases
}

func main() {
	sessioncookie := os.Args[1]
	text := util.GetPuzzleInput("1", sessioncookie)
	input := []int{}
	for _, v := range text {
		num, _ := strconv.Atoi(v)
		input = append(input, num)
	}

	inc := countIncreasingMeasurements(input)
	println("Increasing Measurements: ", inc)
	inc = countSlidingWindowIncreases(input)
	println("Sliding Window: ", inc)
}
