package p2

import (
	"fmt"
  "slices"
  "os"

	"github.com/drcrees/aoc/helpers"
)

type Tile struct {
  x, y int
  adjacent []*Tile
}

func Solve() {
	fmt.Println("--- 9-2 ---")

  strs := helpers.ReadLines("./2025/09/p2/input")
  tiles, grid := tiles(strs)
  // fmt.Println(tiles)
  result := maxarea(tiles, grid)
  // result := 0
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
  minX := min(t1.x, t2.x)
  maxX := max(t1.x, t2.x)
  minY := min(t1.y, t2.y)
  maxY := max(t1.y, t2.y)

  if grid[minY][minX] == 3 { return false }
  if grid[minY][maxX] == 3 { return false }
  if grid[maxY][minX] == 3 { return false }
  if grid[maxY][maxX] == 3 { return false }

  // for y := minY; y <= maxY; y++ {
  //   for x := minX; x <= maxX; x++ {
  //     // fmt.Printf("(%d, %d)\n", x, y)
  //     if grid[y][x] == 3 {
  //       return false
  //     }
  //   }
  // }
  return true
}

func area(t1, t2 *Tile) int {
  return (abs(t2.x - t1.x)+1) * (abs(t2.y - t1.y)+1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func tiles(strs []string) (tiles []*Tile, grid [][]int) {
  maxX, maxY := 0, 0
  for _, str := range strs {
    var x, y int
    fmt.Sscanf(str, "%d,%d", &x, &y)
    tiles = append(tiles, &Tile{x, y, []*Tile{}})
    if x > maxX {
      maxX = x
    }
    if y > maxY {
      maxY = y
    }
  }
  fmt.Printf("Max X: %d, Max Y: %d\n", maxX, maxY)

  grid = make([][]int, maxY+3)
  for i := range grid {
    grid[i] = make([]int, maxX+3)
  }

  for i, t := range tiles {
    if i == len(tiles)-1 {
      t.adjacent = append(t.adjacent, tiles[0])
      t.adjacent = append(t.adjacent, tiles[i-1])
      drawBounds(grid, t.x, t.y, tiles[0].x, tiles[0].y)
      drawBounds(grid, t.x, t.y, tiles[i-1].x, tiles[i-1].y)
      continue
    }
    if i == 0 {
      t.adjacent = append(t.adjacent, tiles[i+1])
      t.adjacent = append(t.adjacent, tiles[len(tiles)-1])
      drawBounds(grid, t.x, t.y, tiles[i+1].x, tiles[i+1].y)
      drawBounds(grid, t.x, t.y, tiles[len(tiles)-1].x, tiles[len(tiles)-1].y)
      continue
    }
    t.adjacent = append(t.adjacent, tiles[i-1])
    t.adjacent = append(t.adjacent, tiles[i+1])
    drawBounds(grid, t.x, t.y, tiles[i-1].x, tiles[i-1].y)
    drawBounds(grid, t.x, t.y, tiles[i+1].x, tiles[i+1].y)
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
  for y := len(grid)-1; y >= 0; y-- {
    for x := len(grid[0])-1; x >= 0; x-- {
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
      // fmt.Printf("%d ", col)
      fmt.Fprintf(f, "%d ", col)
    }
    fmt.Fprintln(f)
    // fmt.Println()
  }
  fmt.Println("Grid printed to file.")
}
