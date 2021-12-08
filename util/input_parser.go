package util

import (
	"strconv"
	"strings"
)

func ParseIntList(input string) []int {
	split := strings.Split(input, ",")
	out := make([]int, len(split))
	for i := range split {
		out[i], _ = strconv.Atoi(split[i])
	}
	return out
}