package p1

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

var directions []Coords = []Coords{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

type Coords struct {
	x, y int
}

func Solve() {
	fmt.Println("--- 4-1 ---")
	grid := helpers.ReadRunes("./2025/04/p1/input")

	result := traverse(grid)

	fmt.Printf("Result: %d\n", result)
}

func traverse(grid [][]rune) int {
	result := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if string(grid[y][x]) == "@" && countAdjacent(grid, x, y) < 4 {
				result++
			}
		}
	}
	return result
}

func countAdjacent(grid [][]rune, x int, y int) int {
	result := 0
	for _, direction := range directions {
		dx, dy := direction.x, direction.y

		if inbounds(grid, x, y, dx, dy) {
			if string(grid[y+dy][x+dx]) == "@" {
				result += 1
			}
		}
	}

	return result
}

func inbounds(grid [][]rune, x int, y int, dx int, dy int) bool {
	if x+dx >= 0 && x+dx < len(grid[0]) {
		return y+dy >= 0 && y+dy < len(grid)
	}

	return false
}
