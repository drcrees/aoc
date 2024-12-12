package p2

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 1-2 ---")
	directions := helpers.ReadFile("./2015/01/p2/input")

	result := 0
	for i, direction := range directions {
		if result == -1 {
			result = i
			break
		}

		if direction == '(' {
			result++
		}
		if direction == ')' {
			result--
		}
	}
	fmt.Printf("Result: %d\n", result)
}
