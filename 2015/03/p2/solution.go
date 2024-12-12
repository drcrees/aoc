package p2

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 3-2 ---")
	directions := helpers.ReadFile("./2015/03/p2/input")

	visited := make(map[[2]int]int)
	visited[[2]int{0, 0}] = 2

	var x1, x2, y1, y2 int
	for i, d := range directions {
		if i%2 == 0 {
			visited = move(&x1, &y1, d, visited)
			continue
		}
		visited = move(&x2, &y2, d, visited)
	}

	result := len(visited)
	fmt.Printf("Result: %d\n", result)
}

func move(x *int, y *int, d rune, visited map[[2]int]int) map[[2]int]int {
	if d == '<' {
		*x--
		visited[[2]int{*x, *y}] += 1
	}
	if d == '>' {
		*x++
		visited[[2]int{*x, *y}] += 1
	}
	if d == '^' {
		*y--
		visited[[2]int{*x, *y}] += 1
	}
	if d == 'v' {
		*y++
		visited[[2]int{*x, *y}] += 1
	}

	return visited
}
