package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"riedmann.dev/aoc21/util"
)

func parseHeightmap(input []string) [][]int {
	heightMap := make([][]int, len(input))
	for i := range heightMap {
		heightMap[i] = make([]int, len(input[0]))
	}

	for y, s := range(input) {
		for  x, c := range(strings.Split(s, "")) {
			i, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalln("Expected only integer input, but got ", c)
			}

			heightMap[y][x] = i
		}
	}

	return heightMap
}

func findLowPoints(hMap [][]int) (lowPoints []int) {
	for y := range(hMap) {
		for x := range(hMap[y]) {
			if smallerThanNeighbours(x, y, hMap) {
				lowPoints = append(lowPoints, hMap[y][x])
			}
		}
	}
	return lowPoints
}

func smallerThanNeighbours(x int, y int, hMap [][]int) bool {
	height := len(hMap)
	width := len(hMap[0])
	//top (y-1)
	if y-1 >= 0 && hMap[y-1][x] <= hMap[y][x] {
		return false
	}
	//right (x+1)
	if x+1 < width && hMap[y][x+1] <= hMap[y][x] {
		return false
	}
	//bottom (y+1)
	if y+1 < height && hMap[y+1][x] <= hMap[y][x] {
		return false
	}
	//left (x-1)
	if x-1 >= 0 && hMap[y][x-1] <= hMap[y][x] {
		return false
	}
	return true
}

func calculateRiskLevels(lowPoints []int) (risk int) {
	for _, p := range(lowPoints) {
		risk += (p + 1)
	}
	return
}

func calculateSolution(input []string) int {
    heightMap := parseHeightmap(input)
	lowPoints := findLowPoints(heightMap)
    return calculateRiskLevels(lowPoints)
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Session Cookie needed as arg to get puzzle input!")
    }
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution: ", calculateSolution(input))
}
const day = "9"
