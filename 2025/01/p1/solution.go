package p1

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 1-1 ---")
	lines := helpers.ReadLines("./2025/01/p1/input")

	result := 0
	current := 50

	for _, line := range lines {
		current = dial(current, line)
		if current%100 == 0 {
			result++
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func dial(current int, line string) int {
	runes := []rune(line)

	modifier := 1
	direction := string(runes[0:1])
	magnitude, _ := strconv.Atoi(string(runes[1:]))

	if direction == "L" {
		modifier = -1
	}

	return current + (modifier * magnitude)
}
