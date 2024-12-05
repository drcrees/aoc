package p1

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 2-1 ---")
	strs := helpers.ReadDelimitedStrings("./2024/02/p1/input", " ")

	result := 0
	for _, numStrs := range strs {
		if IsSafe(numStrs, 0, 0, false) {
			result++
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func IsSafe(numStrs []string, index int, dir int, isSafe bool) bool {
	if index == len(numStrs)-1 {
		return isSafe
	}

	n1, _ := strconv.Atoi(numStrs[index])
	n2, _ := strconv.Atoi(numStrs[index+1])

	diff := n1 - n2
	if diff == 0 || diff < -3 || diff > 3 {
		return false
	}

	// change in direction
	if (dir < 0 && diff > 0) || (dir > 0 && diff < 0) {
		return false
	}

	return IsSafe(numStrs, index+1, diff, true)
}
