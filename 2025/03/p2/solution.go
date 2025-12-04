package p2

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 3-2 ---")
	lines := helpers.ReadLines("./2025/03/p2/input")

	result := 0

	for _, line := range lines {
		digits := parseLine(line)

		result += largest(digits, 12)
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

func largest(a []int, b int) int {
	result := 0
	start, end := 0, len(a)-b+1

	for i := 0; i < b; i++ {
		slice := a[start:end]

		max, index := 0, -1
		for j, num := range slice {
			if num > max {
				max = num
				index = j
			}
		}

		result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, max))
		start += index + 1
		end += 1
	}
	return result
}
