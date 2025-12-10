package p2

import (
	"cmp"
	"fmt"
	"os"
	"slices"

	"github.com/drcrees/aoc/helpers"
)

type Tile struct {
	x, y, cx, cy int
	adjacent     []*Tile
}

func Solve() {
	fmt.Println("--- 9-2 ---")

	strs := helpers.ReadLines("./2025/09/p2/input")
	tiles, grid := tiles(strs)
	result := maxarea(tiles, grid)
	fmt.Printf("Result: %d\n", result)
}

func maxarea(tiles []*Tile, grid [][]int) int {
	areas := []int{}
	for _, t1 := range tiles {
		for _, t2 := range tiles {
			if t1 != t2 {
				if valid(t1, t2, grid) {
					areas = append(areas, area(t1, t2))
				}
			}
		}
	}
	slices.Sort(areas)
	return areas[len(areas)-1]
}

func valid(t1, t2 *Tile, grid [][]int) bool {
	minX := min(t1.cx, t2.cx)
	maxX := max(t1.cx, t2.cx)
	minY := min(t1.cy, t2.cy)
	maxY := max(t1.cy, t2.cy)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y][x] == 3 {
				return false
			}
		}
	}
	return true
}

func area(t1, t2 *Tile) int {
	return (abs(t2.x-t1.x) + 1) * (abs(t2.y-t1.y) + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func tiles(strs []string) (tiles []*Tile, grid [][]int) {
	for _, str := range strs {
		var x, y int
		fmt.Sscanf(str, "%d,%d", &x, &y)
		tiles = append(tiles, &Tile{x, y, 0, 0, []*Tile{}})
	}

	for i, t := range tiles {
		if i == len(tiles)-1 {
			t.adjacent = append(t.adjacent, tiles[0])
			t.adjacent = append(t.adjacent, tiles[i-1])
			continue
		}
		if i == 0 {
			t.adjacent = append(t.adjacent, tiles[i+1])
			t.adjacent = append(t.adjacent, tiles[len(tiles)-1])
			continue
		}
		t.adjacent = append(t.adjacent, tiles[i-1])
		t.adjacent = append(t.adjacent, tiles[i+1])
	}

	slices.SortFunc(tiles, func(a, b *Tile) int {
		return cmp.Compare(a.x, b.x)
	})

	i := 0
	x := tiles[0].x
	for _, tile := range tiles {
		if tile.x == x {
			tile.cx = i
		} else {
			i++
			x = tile.x
			tile.cx = i
		}
	}

	slices.SortFunc(tiles, func(a, b *Tile) int {
		return cmp.Compare(a.y, b.y)
	})

	i = 0
	y := tiles[0].y
	for _, tile := range tiles {
		if tile.y == y {
			tile.cy = i
		} else {
			i++
			y = tile.y
			tile.cy = i
		}
	}

	grid = make([][]int, len(tiles))
	for i := range grid {
		grid[i] = make([]int, len(tiles))
	}

	for _, t := range tiles {
		for _, adj := range t.adjacent {
			drawBounds(grid, t.cx, t.cy, adj.cx, adj.cy)
		}
	}

	flood(grid)
	printGrid(grid)
	return tiles, grid
}

func flood(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 2 {
				break
			}
			grid[y][x] = 3
		}
	}
	for y := len(grid) - 1; y >= 0; y-- {
		for x := len(grid[0]) - 1; x >= 0; x-- {
			if grid[y][x] == 2 {
				break
			}
			grid[y][x] = 3
		}
	}
}

func drawBounds(grid [][]int, fromX, fromY, toX, toY int) {
	for y := fromY; y <= toY; y++ {
		for x := fromX; x <= toX; x++ {
			grid[y][x] = 2
		}
	}
}

func printGrid(grid [][]int) {
	f, _ := os.Create("grid")
	defer f.Close()

	for _, row := range grid {
		for _, col := range row {
			fmt.Fprintf(f, "%d ", col)
		}
		fmt.Fprintln(f)
	}
}
