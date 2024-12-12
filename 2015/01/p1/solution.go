package p1

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 1-1 ---")
	directions := helpers.ReadFile("./2015/01/p1/input")

	result := 0
	for _, direction := range directions {
		if direction == '(' {
			result++
		}
		if direction == ')' {
			result--
		}
	}
	fmt.Printf("Result: %d\n", result)
}
