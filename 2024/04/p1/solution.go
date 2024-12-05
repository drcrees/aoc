package p1

import (
	"fmt"
	"github.com/drcrees/aoc/helpers"
)

type Coords struct {
	x int
	y int
}

func directions() map[int]Coords {
	return map[int]Coords{
		0: {0, 0}, 1: {0, 1}, 2: {0, -1}, 3: {1, 0}, 4: {-1, 0}, 5: {1, 1}, 6: {1, -1}, 7: {-1, 1}, 8: {-1, -1},
	}
}

func Solve() {
	fmt.Println("--- 4-1 ---")
	grid := helpers.ReadRunes("./2024/04/p1/input")

	result := whereXmas(grid)
	fmt.Printf("Result: %d\n", result)
}

func whereXmas(grid [][]rune) (result int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'X' {
				result += isXmas(grid, x, y, 'M', 0)
			}
		}
	}

	return result
}

func inbounds(start Coords, op Coords, bounds Coords) bool {
	if start.x+op.x >= 0 && start.x+op.x <= bounds.x {
		return start.y+op.y >= 0 && start.y+op.y <= bounds.y
	}

	return false
}

func isXmas(grid [][]rune, x int, y int, letter rune, direction int) (numXmas int) {
	for key, dir := range directions() {
		if inbounds(Coords{x, y}, dir, Coords{len(grid) - 1, len(grid[0]) - 1}) {
			if grid[y+dir.y][x+dir.x] == letter {
				if direction == 0 {
					numXmas += isXmas(grid, x+dir.x, y+dir.y, 'A', key)
				}
				if key == direction {
					if letter == 'S' {
						return 1
					}
					numXmas += isXmas(grid, x+dir.x, y+dir.y, 'S', key)
				}
			}
		}
	}

	return numXmas
}
