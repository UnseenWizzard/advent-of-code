package util

import (
	"strconv"
	"strings"
)

func ParseIntMap(input []string) [][]int {
	out := make([][]int, len(input))
	for i, l := range(input) {
		out[i] = ParseIntListBy(l, "")

	}
	return out
}

func ParseIntList(input string) []int {
	return ParseIntListBy(input, ",")
}

func ParseIntListBy(input string, sep string) []int {
	split := strings.Split(input, sep)
	out := make([]int, len(split))
	for i := range split {
		out[i], _ = strconv.Atoi(split[i])
	}
	return out
}