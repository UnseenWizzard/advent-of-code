package main

import (
	"log"
	"os"
	"riedmann.dev/aoc21/util"
	"strconv"
	"strings"
)

func applyCommand(cmd string, pos int, depth int, aim int) (int, int, int) {
	if len(cmd) == 0 {
		println("Ignoring empty input")
		return pos, depth, aim
	}
	cmdSplit := strings.Split(cmd, " ")
	action := cmdSplit[0]
	steps, err := strconv.Atoi(cmdSplit[1])
	if err != nil {
		log.Fatalln("Expected a number but got ", cmdSplit[1])
	}
	switch action {
	case "forward":
		return pos + steps, depth + (aim * steps), aim
	case "up":
		return pos, depth, aim - steps
	case "down":
		return pos, depth, aim + steps
	}
	return pos, depth, aim
}

func applyCommands(cmds []string) (int, int) {
	pos := 0
	depth := 0
	aim := 0
	for _, c := range cmds {
		pos, depth, aim = applyCommand(c, pos, depth, aim)
	}
	return pos, depth
}

func calculateSolution(pos int, depth int) int {
	return pos * depth
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Session Cookie needed as arg to get puzzle input!")
	}
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput("2", sessioncookie)

	finalPos, finalDepth := applyCommands(input)
	println("Result:", calculateSolution(finalPos, finalDepth))
}
