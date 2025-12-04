package p1

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 3-1 ---")
	lines := helpers.ReadLines("./2025/03/p1/input")

	result := 0

	for _, line := range lines {
		digits := parseLine(line)

		_, combo := largest(0, digits)
		result += combo
	}
	fmt.Printf("Result: %d\n", result)
}

func parseLine(str string) []int {
	var digits []int
	for _, digit := range str {
		digits = append(digits, int(digit)-'0')
	}
	return digits
}

func largest(a int, b []int) (int, int) {
	if len(b) == 1 {
		return b[0], b[0]
	}

	largest, next := largest(b[0], b[1:])
	current, _ := strconv.Atoi(fmt.Sprintf("%d%d", b[0], largest))
	if current >= next {
		if b[0] < largest {
			return largest, current
		}
		return b[0], current
	}

	if b[0] < largest {
		return largest, next
	}
	return b[0], next
}
