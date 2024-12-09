package p1

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

type Coords struct {
  x, y int
}

type Antenna struct {
  frequency rune
  coords Coords
}

type Antinode struct {
  coords Coords
}

func Solve() {
	fmt.Println("--- 8-1 ---")
  grid := helpers.ReadRunes("./2024/08/p1/input")

  antennas := findAntennas(grid)

  result := 0
  for f, a := range antennas {
    for _, a1 := range a {
      for _, a2 := range a {
        an1, an2 := antinodes(a1.coords, a2.coords)
        if inbounds(an1.x, an1.y, len(grid[0]), len(grid)) {
          if grid[an1.y][an1.x] != f && grid[an1.y][an1.x] != '#' {
            grid[an1.y][an1.x] = '#'
            result++
          }
        }
        if inbounds(an2.x, an2.y, len(grid[0]), len(grid)) {
          if grid[an2.y][an2.x] != f && grid[an2.y][an2.x] != '#' {
            grid[an2.y][an2.x] = '#'
            result++
          }
        }
      }
    }
  }

	fmt.Printf("Result: %d\n", result)
}

func antinodes(from Coords, to Coords) (Coords, Coords) {
  delta := Coords{to.x-from.x, to.y-from.y}

  a1 := Coords{delta.x*-1 + from.x, delta.y*-1 + from.y}
  a2 := Coords{delta.x+to.x, delta.y+to.y}

  return a1, a2
}

func inbounds(x int, y int, maxX int, maxY int) bool {
	if x >= 0 && x < maxX {
		return y >= 0 && y < maxY
	}

	return false
}

func findAntennas(grid [][]rune) map[rune][]*Antenna{
  abf := make(map[rune][]*Antenna)
  for y := 0; y < len(grid); y++ {
    for x := 0; x < len(grid[y]); x++ {
      if grid[y][x] != '.' {
        antenna := &Antenna{
          frequency: grid[y][x],
          coords: Coords{x,y},
        }
        if abf[grid[y][x]] != nil {
          abf[grid[y][x]] = append(abf[grid[y][x]], antenna)
          continue
        }
        abf[grid[y][x]] = []*Antenna{antenna}
      }
    }
  }

  return abf
}

func printGrid(grid [][]rune) {
  for y := 0; y < len(grid); y++ {
    for x := 0; x < len(grid[y]); x++ {
      fmt.Printf("%c ", grid[y][x])
    }
    fmt.Println()
  }
}
