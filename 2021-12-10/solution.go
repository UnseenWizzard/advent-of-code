package main

import (
	"container/list"
	"log"
	"os"
	"strings"

	"riedmann.dev/aoc21/util"
)

var pointMap = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197, 
	">": 25137,
}


func isOpeningChar(c string) bool {
	if c == "(" || c == "[" || c == "{" || c == "<" {
		return true
	}
	return false
}

func isExpectedClosing(c string, open string) bool {
	if open == "(" && c == ")" {
		return true
	} 
	if open == "{" && c == "}" {
		return true
	} 
	if open == "[" && c == "]" {
		return true
	} 
	if open == "<" && c == ">" {
		return true
	} 
	return false
}

func calculateSolution(input []string) int {
	points := 0
	for _, l := range(input) {
		toClose := list.New()
		for _, c := range(strings.Split(l, "")) {
			if isOpeningChar(c) {
				toClose.PushBack(c)
			} else if toClose.Len() >0 && isExpectedClosing(c, toClose.Back().Value.(string)) {
				toClose.Remove(toClose.Back())
			} else {
				println("failed on ", c)
				points += pointMap[c]
				break
			}
		}
	}
    return points
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Session Cookie needed as arg to get puzzle input!")
    }
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input))
}
const day = "10"
