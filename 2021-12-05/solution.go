package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"riedmann.dev/aoc21/util"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

type ventMap [][]int

func parseInput(input []string) (lines []line, xDimension int, yDimension int) {
	lines = make([]line, len(input))
	for i, l := range input {
		s := strings.Split(l, " -> ")
		if len(s) != 2 {
			log.Fatal("failed to parse input ", l)
		}
		start := parsePoint(s[0])
		end := parsePoint(s[1])

		if start.x > xDimension {
			xDimension = start.x
		}
		if end.x > xDimension {
			xDimension = end.x
		}
		if start.y > yDimension {
			yDimension = start.y
		}
		if end.y > yDimension {
			yDimension = end.y
		}

		lines[i] = line{start, end}
	}
	return
}

func parsePoint(input string) point {
	p := strings.Split(input, ",")
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	return point{x, y}
}

func createMap(lines []line, xDimension int, yDimension int, verticalMapping bool) ventMap {
	vMap := make(ventMap, yDimension+1)
	for i := range vMap {
		vMap[i] = make([]int, xDimension+1)
	}

	for _, l := range lines {
		if l.start.x == l.end.x {
			//vertical line
			start, end := sortDirection(l.start.y, l.end.y)
			for i := start; i <= end; i++ {
				vMap[i][l.start.x] = vMap[i][l.start.x] + 1
			}
		} else if l.start.y == l.end.y {
			//horizontal line
			start, end := sortDirection(l.start.x, l.end.x)
			for i := start; i <= end; i++ {
				vMap[l.start.y][i] = vMap[l.start.y][i] + 1
			}
		} else if verticalMapping {
			y := l.start.y
			for x := l.start.x; x != l.end.x; x = advance(x, l.start.x, l.end.x) {
				vMap[y][x] = vMap[y][x] + 1
				y = advance(y, l.start.y, l.end.y)
			}
			vMap[l.end.y][l.end.x] = vMap[l.end.y][l.end.x] + 1
		} else {
			println("Ignoring diagonal line")
		}
	}

	return vMap
}

func advance(pos int, start int, end int) int {
	if start > end {
		return pos-1
	}
	if start < end {
		return pos+1
	}
	return pos
}

func sortDirection(a int, b int) (start int, end int) {
	if a >= b {
		return b, a
	}
	return a, b
}

func countIntersectingLines(vMap ventMap) (intersections int) {
	for _, r := range vMap {
		for _, c := range r {
			if c > 1 {
				intersections++
			}
		}
	}
	return
}

func calculateSolution(input []string, verticalMapping bool) int {
	lines, xDimension, yDimension := parseInput(input)
	vMap := createMap(lines, xDimension, yDimension, verticalMapping)
	count := countIntersectingLines(vMap)
	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Session Cookie needed as arg to get puzzle input!")
	}
	sessioncookie := os.Args[1]
	input := util.GetPuzzleInput(day, sessioncookie)

	println("Solution  : ", calculateSolution(input, false))
	println("Part 2 Sol: ", calculateSolution(input, true))
}

const day = "5"
