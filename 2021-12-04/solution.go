package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"riedmann.dev/aoc21/util"
)

type cell struct {
	val    int
	marked bool
}

type board [][]cell

func parseInput(input []string) (numbers []int, boards []board) {
	if len(input) < 3 {
		log.Println("Input does not match expectation:", input)
		return
	}
	for _, c := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal("First input row contained non-integer!")
		}
		numbers = append(numbers, i)
	}

	boardIndex := -1
	rowIndex := 0
	for _, l := range input[1:] {
		if len(l) == 0 {
			b := make(board, 5)
			for i := range b {
				b[i] = make([]cell, 5)
			}
			boards = append(boards, b)
			boardIndex++
			rowIndex = 0
			continue
		}
		for cellIndex, char := range strings.Fields(l) {
			num, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal("Board contained non-integer value!")
			}
			boards[boardIndex][rowIndex][cellIndex] = cell{num, false}
		}
		rowIndex++
	}

	return
}

func markBoard(b board, val int) (won bool) {
	winningCols := make([]bool, len(b[0]))
	for i := range winningCols {
		winningCols[i] = true
	}

	for rI := 0; rI < len(b); rI++ {
		winningRow := true
		for cI := 0; cI < len(b[rI]); cI++ {
			if b[rI][cI].val == val {
				b[rI][cI].marked = true
			}
			winningRow = winningRow && b[rI][cI].marked
			winningCols[cI] = winningCols[cI] && b[rI][cI].marked
		}
		if winningRow {
			return true
		}
	}
	for _, winningCol := range winningCols {
		if winningCol {
			return true
		}
	}
	return false
}

func calcuateBoardScore(b board, winningNum int) int {
	sumUnmarked := 0
	for _, row := range b {
		for _, c := range row {
			if !c.marked {
				sumUnmarked = sumUnmarked + c.val
			}
		}
	}
	return sumUnmarked * winningNum
}

func calculateSolution(input []string) (firstWinningScore int, lastWinningScore int) {
	numbers, boards := parseInput(input)

	boardWon := make([]bool, len(boards))

	for _, num := range numbers {
		for bI, b := range boards {
			won := markBoard(b, num)
			if won {
				/* For the initial puzzle, we can early exit and return the first winning score.
				 * To also solve part 2, we run through all numbers,
				 * so we know the last winning board's score in addition to the first one to win.
				 */

				if !boardWon[bI] {
					boardWon[bI] = true
					lastWinningScore = calcuateBoardScore(b, num)
					if firstWinningScore <= 0 {
						firstWinningScore = lastWinningScore
					}
				}
			}
		}
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Session Cookie needed as arg to get puzzle input!")
	}
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	firstWinningScore, lastWinningScore := calculateSolution(input)
	println("Solution #1: ", firstWinningScore)
	println("Solution #2: ", lastWinningScore)
}

const day = "4"
