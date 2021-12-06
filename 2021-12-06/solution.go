package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"riedmann.dev/aoc21/util"
)

func passDay(fish []int) []int {
	numFish := len(fish)
	for i := 0; i < numFish; i++ {
		if fish[i] == 0 {
			fish[i] = 6
			fish = append(fish, 8)
		} else {
			fish[i]--
		}
	}
	return fish
}

func passDays(fish []int, days int) int {
	for d := 0; d < days; d++ {
		fish = passDay(fish)
	}
	return len(fish)
}

func calculateSolution(input []string, days int) int {
	if len(input) != 1 {
		log.Fatalln("Expected single line of input but got ", len(input))
	}

	fishString := strings.Split(input[0], ",")
	fish := make([]int, len(fishString))
	for i := range fishString {
		fish[i], _ = strconv.Atoi(fishString[i])
	}


	numFish := passDays(fish, days)
    return numFish
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Session Cookie needed as arg to get puzzle input!")
    }
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input, 80))
}
const day = "6"
