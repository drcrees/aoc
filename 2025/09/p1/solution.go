package p1

import (
	"fmt"
	"slices"

	"github.com/drcrees/aoc/helpers"
)

type Tile struct {
	x, y int
}

func Solve() {
	fmt.Println("--- 9-1 ---")

	strs := helpers.ReadLines("./2025/09/p1/input")
	tiles := tiles(strs)
	result := maxarea(tiles)

	fmt.Printf("Result: %d\n", result)
}

func maxarea(tiles []Tile) int {
	areas := []int{}
	for _, t1 := range tiles {
		for _, t2 := range tiles {
			if t1 != t2 {
				areas = append(areas, area(t1, t2))
			}
		}
	}
	slices.Sort(areas)
	return areas[len(areas)-1]
}

func area(t1, t2 Tile) int {
	return (abs(t2.x-t1.x) + 1) * (abs(t2.y-t1.y) + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func tiles(strs []string) (tiles []Tile) {
	for _, str := range strs {
		var x, y int
		fmt.Sscanf(str, "%d,%d", &x, &y)
		tiles = append(tiles, Tile{x, y})
	}
	return tiles
}
