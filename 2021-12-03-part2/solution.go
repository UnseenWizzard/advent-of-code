package main

import (
	"log"
	"os"
	"riedmann.dev/aoc21/util"
	"strconv"
	"strings"
)

type ratingType int

const (
	oxygenGenerator ratingType = iota
	co2Scrubber                = iota
)

func getBit(input string, pos int) int {
	c := strings.Split(input, "")[pos]
	bit, _ := strconv.Atoi(c)
	return bit
}

func findCriterion(input []string, pos int, ratingType ratingType) (bitCriterion string) {
	if len(input) < 1 {
		log.Fatal("Can't find most common bit for empty input")
	}

	middle := float64(len(input)) / 2.0
	sum := 0
	for _, s := range input {
		bit := getBit(s, pos)
		sum += bit
	}

	if float64(sum) >= middle {
		if ratingType == oxygenGenerator {
			return "1"
		}
		return "0"
	} else {
		if ratingType == oxygenGenerator {
			return "0"
		}
		return "1"
	}
}

func findRating(input []string, pos int, ratingType ratingType) string {
	if len(input) < 1 {
		log.Fatal("Can't find most common bit for empty input")
	}

	var nextToConsider []string

	bitCriterion := findCriterion(input, pos, ratingType)
	for _, s := range input {
		c := strings.Split(s, "")[pos]
		if c == bitCriterion {
			nextToConsider = append(nextToConsider, s)
		}
	}

	if len(nextToConsider) == 1 {
		return nextToConsider[0]
	}

	pos++
	return findRating(nextToConsider, pos, ratingType)
}

func calculateSolution(input []string) int {
	if len(input) < 1 {
		return 0
	}
	oxygenString := findRating(input, 0, oxygenGenerator)
	co2String := findRating(input, 0, co2Scrubber)

	println(oxygenString)
	println(co2String)

	oxygen, _ := strconv.ParseInt(oxygenString, 2, 64)
	co2, _ := strconv.ParseInt(co2String, 2, 64)
	lifesupportRating := oxygen * co2
	return int(lifesupportRating)
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
