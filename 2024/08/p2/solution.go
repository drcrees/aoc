package p2

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
  delta Coords
}

func Solve() {
	fmt.Println("--- 8-2 ---")
  grid := helpers.ReadRunes("./2024/08/p2/input")

  antennas := findAntennas(grid)

  result := 0
  for _, a := range antennas {
    for i, a1 := range a {
      result++
      for j, a2 := range a {
        if i != j {
          an1, an2 := antinodes(a1.coords, a2.coords)
          dx := 0
          dy := 0
          for inbounds(an1.coords.x+dx, an1.coords.y+dy, len(grid[0]), len(grid)) {
            if grid[an1.coords.y+dy][an1.coords.x+dx] == '.' && grid[an1.coords.y+dy][an1.coords.x+dx] != '#' {
              grid[an1.coords.y+dy][an1.coords.x+dx] = '#'
              result++
            }
            dx += an2.delta.x
            dy += an2.delta.y
          }
          dx = 0
          dy = 0
          for inbounds(an2.coords.x+dx, an2.coords.y+dy, len(grid[0]), len(grid)) {
            if grid[an2.coords.y+dy][an2.coords.x+dx] == '.' && grid[an2.coords.y+dy][an2.coords.x+dx] != '#' {
              grid[an2.coords.y+dy][an2.coords.x+dx] = '#'
              result++
            }
            dx += an2.delta.x
            dy += an2.delta.y
          }
        }
      }
    }
  }

	fmt.Printf("Result: %d\n", result)
}

func antinodes(from Coords, to Coords) (Antinode, Antinode) {
  delta := Coords{to.x-from.x, to.y-from.y}

  a1 := Antinode{Coords{delta.x*-1 + from.x, delta.y*-1 + from.y}, Coords{delta.x*-1, delta.y*-1}}
  a2 := Antinode{Coords{delta.x+to.x, delta.y+to.y}, Coords{delta.x, delta.y}}

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
