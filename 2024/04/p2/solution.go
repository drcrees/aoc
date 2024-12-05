package p2

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

type Coords struct {
	x int
	y int
}

var xmap map[Coords]int

func directions() map[int]Coords {
	return map[int]Coords{
		0: {0, 0}, 1: {1, 1}, 2: {1, -1}, 3: {-1, 1}, 4: {-1, -1},
	}
}

func Solve() {
	fmt.Println("--- 4-2 ---")
	grid := helpers.ReadRunes("./2024/04/p2/input")

	result := whereXmas(grid)
	fmt.Printf("Result: %d\n", result)
}

func whereXmas(grid [][]rune) (result int) {
	xmap := make(map[Coords]int)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'M' {
				isXmas(grid, xmap, x, y, 'A', 0)
			}
		}
	}

	for _, v := range xmap {
		if v >= 2 {
			result++
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

func isXmas(grid [][]rune, xmap map[Coords]int, x int, y int, letter rune, direction int) bool {
	for key, dir := range directions() {
		if inbounds(Coords{x, y}, dir, Coords{len(grid) - 1, len(grid[0]) - 1}) {
			if grid[y+dir.y][x+dir.x] == letter {
				if direction == 0 && isXmas(grid, xmap, x+dir.x, y+dir.y, 'S', key) {
					xmap[Coords{x + dir.x, y + dir.y}]++
					continue
				}
				if key == direction {
					if letter == 'S' {
						return true
					}
				}
			}
		}
	}

	return false
}
