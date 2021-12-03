package main

import (
	"log"
	"os"
	"riedmann.dev/aoc21/util"
	"strconv"
	"strings"
)

func calculateSolution(input []string) int {
	if len(input) < 1 {
		return 0
	}
	epsilonString := ""
	gammaString := ""

	numValues := len(input)

	for i := 0; i < len(input[0]); i++ {
		sum := 0
		for _, s := range input {
			b := strings.Split(s, "")[i]
			bit, _ := strconv.Atoi(b)
			sum += bit
		}
		println(sum)
		if sum > numValues/2 {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}

	println(epsilonString)
	println(gammaString)

	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)
	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	powerConsumption := epsilon * gamma
	return int(powerConsumption)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Session Cookie needed as arg to get puzzle input!")
	}
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input[:len(input)-1]))
}

const day = "3"
