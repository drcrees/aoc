package p1

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

type Coords struct {
	x, y int
}

func Solve() {
	fmt.Println("--- 10-1 ---")

	grid := helpers.ReadInts("./2024/10/p1/input")

	trailheads := findTrailheads(grid)

	result := 0
	for _, th := range trailheads {
		result += hike(th, grid, make(map[Coords]int))
	}
	fmt.Printf("Result: %d\n", result)
}

func hike(coords Coords, grid [][]int, found map[Coords]int) int {
	result := 0
	current := grid[coords.y][coords.x]
	if grid[coords.y][coords.x] == 9 {
		if _, ok := found[Coords{coords.x, coords.y}]; ok {
			return 0
		}
		found[Coords{coords.x, coords.y}] = 1
		return 1
	}

	if inbounds(coords.x, coords.y, -1, 0, grid) && grid[coords.y][coords.x-1]-current == 1 {
		result += hike(Coords{coords.x - 1, coords.y}, grid, found)
	}

	if inbounds(coords.x, coords.y, 1, 0, grid) && grid[coords.y][coords.x+1]-current == 1 {
		result += hike(Coords{coords.x + 1, coords.y}, grid, found)
	}

	if inbounds(coords.x, coords.y, 0, -1, grid) && grid[coords.y-1][coords.x]-current == 1 {
		result += hike(Coords{coords.x, coords.y - 1}, grid, found)
	}

	if inbounds(coords.x, coords.y, 0, 1, grid) && grid[coords.y+1][coords.x]-current == 1 {
		result += hike(Coords{coords.x, coords.y + 1}, grid, found)
	}

	return result
}

func inbounds(x, y, dx, dy int, grid [][]int) bool {
	if x+dx >= 0 && x+dx < len(grid[0]) {
		return y+dy >= 0 && y+dy < len(grid)
	}

	return false
}
func findTrailheads(grid [][]int) []Coords {
	var trailheads []Coords
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				trailheads = append(trailheads, Coords{x, y})
			}
		}
	}
	return trailheads
}

func printGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			fmt.Printf("%d ", grid[y][x])
		}
		fmt.Println()
	}
}
