package p1

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 3-1 ---")
	directions := helpers.ReadFile("./2015/03/p1/input")

	visited := make(map[[2]int]int)
	visited[[2]int{0, 0}] = 1

	var x, y int
	for _, d := range directions {
		if d == '<' {
			x--
			visited[[2]int{x, y}] += 1
		}
		if d == '>' {
			x++
			visited[[2]int{x, y}] += 1
		}
		if d == '^' {
			y--
			visited[[2]int{x, y}] += 1
		}
		if d == 'v' {
			y++
			visited[[2]int{x, y}] += 1
		}
	}

	result := len(visited)
	fmt.Printf("Result: %d\n", result)
}
