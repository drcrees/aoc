package p2

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 1-2 ---")
	lines := helpers.ReadLines("./2025/01/p2/input")
	result := 0
	current := 50

	for _, line := range lines {
		position, value := dial(current, line)
		current = position
		result += value
	}

	fmt.Printf("Result: %d\n", result)
}

func dial(position int, line string) (int, int) {
	result := 0
	runes := []rune(line)

	direction := string(runes[0:1])
	magnitude, _ := strconv.Atoi(string(runes[1:]))

	left := position - (magnitude % 100)
	right := position + (magnitude % 100)

	rotations := magnitude / 100
	result += rotations

	if direction == "R" {
		if right > 100 {
			result += 1
		}
		position = right % 100
	}

	if direction == "L" {
		if magnitude%100 > position {
			if position != 0 {
				result += 1
			}
			position = left + 100
		} else {
			position = left
		}
	}

	if position == 0 {
		result += 1
	}

	return position, result
}
