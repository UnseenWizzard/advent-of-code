package main

import (
	"os"
	"log"
	"riedmann.dev/aoc21/util"
)

func increment(octopi [][]int) {
	for i := range(octopi) {
		for j := range(octopi[i]) {
			octopi[i][j] += 1
		}
	}
}

func flash(i int, j int, octopi [][]int, alreadyFlashed [][]bool) {
	if alreadyFlashed[i][j] {
		return
	} 
	alreadyFlashed[i][j] = true

	height := len(octopi)
	width := len(octopi[0])
	//top neighbours
	if i-1 >= 0 {
		if j-1 >= 0 {
			octopi[i-1][j-1] += 1
			if octopi[i-1][j-1] > 9 {
				flash(i-1, j-1, octopi, alreadyFlashed)
			}
		}
		octopi[i-1][j] += 1
		if octopi[i-1][j] > 9 {
			flash(i-1, j, octopi, alreadyFlashed)
		}
		if j+1 < width {
			octopi[i-1][j+1] += 1
			if octopi[i-1][j+1] > 9 {
				flash(i-1, j+1, octopi, alreadyFlashed)
			}
		}
	}
	//right
	if j+1 < width {
		octopi[i][j+1] += 1
		if octopi[i][j+1] > 9 {
			flash(i, j+1, octopi, alreadyFlashed)
		}
	}
	//left
	if j-1 >= 0 {
		octopi[i][j-1] += 1
		if octopi[i][j-1] > 9 {
			flash(i, j-1, octopi, alreadyFlashed)
		}
	}
	//bottom neighbours
	if i+1 < height {
		if j-1 >= 0 {
			octopi[i+1][j-1] += 1
			if octopi[i+1][j-1] > 9 {
				flash(i+1, j-1, octopi, alreadyFlashed)
			}
		}
		octopi[i+1][j] += 1
		if octopi[i+1][j] > 9 {
			flash(i+1, j, octopi, alreadyFlashed)
		}
		if j+1 < width {
			octopi[i+1][j+1] += 1
			if octopi[i+1][j+1] > 9 {
				flash(i+1, j+1, octopi, alreadyFlashed)
			}
		}
	}
}

func step(octopi [][]int) int {
	increment(octopi)

	flashed := make([][]bool, len(octopi))
	for i := range(flashed) {
		flashed[i] = make([]bool, len(octopi[0]))
	}

	for i := range(octopi) {
		for j := range(octopi[i]) {
			if octopi[i][j] > 9 {
				flash(i,j,octopi, flashed)
			}
		}
	}

	count := 0
	for i := range(flashed) {
		for j := range(flashed[i]) { 
			if flashed[i][j] {
				octopi[i][j] = 0
				count++
			}
		}
	}
	return count
}

func calculateSolution(input []string, steps int) (flashed int) {
    octopi := util.ParseIntMap(input)
	for i := 0; i < steps; i++ {
		flashed += step(octopi)
	}
    return 
}

func calculatePartTwo(input []string) (allFlashedInStep int) {
    octopi := util.ParseIntMap(input)
	steps := 0
	for {
		steps++
		flashed := step(octopi)
		if flashed == (len(octopi)) * (len(octopi[0])) {
			return steps
		}
	}
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Session Cookie needed as arg to get puzzle input!")
    }
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input, 100))
	println("All flashed in step: ", calculatePartTwo(input))
}
const day = "11"
