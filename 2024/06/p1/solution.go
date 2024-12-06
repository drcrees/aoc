package p1

import (
	"fmt"

	"github.com/drcrees/aoc/helpers"
)

var visits = 0

type Map struct {
	guard *Guard
	tiles [][]*Tile
}

type Tile struct {
	x, y       int
	isObstacle bool
	visited    bool
}

type Guard struct {
	x, y, dx, dy int
	tilesVisited []*Tile
	m            *Map
}

func Solve() {
	fmt.Println("--- 6-1 ---")
	grid := helpers.ReadRunes("./2024/06/p1/input")

	m := buildMap(grid)
	m.guard.Patrol()

	result := len(m.guard.tilesVisited)
	fmt.Printf("Result: %d\n", result)
}

func (g *Guard) Patrol() {
	for g.inbounds() {
		if g.m.tiles[g.y+g.dy][g.x+g.dx].isObstacle {
			t := g.dx
			g.dx = g.dy * -1
			g.dy = t
			continue
		}

		g.x = g.x + g.dx
		g.y = g.y + g.dy

		if !g.m.tiles[g.y][g.x].visited {
			g.m.tiles[g.y][g.x].visited = true
			g.tilesVisited = append(g.tilesVisited, g.m.tiles[g.y][g.x])
		}
	}
}

func (g *Guard) inbounds() bool {
	if g.x+g.dx >= 0 && g.x+g.dx < len(g.m.tiles[0]) {
		return g.y+g.dy >= 0 && g.y+g.dy < len(g.m.tiles)
	}

	return false
}

func buildMap(grid [][]rune) *Map {
	m := Map{
		guard: nil,
		tiles: make([][]*Tile, 0),
	}

	for y := 0; y < len(grid); y++ {
		var row []*Tile
		for x := 0; x < len(grid[0]); x++ {
			tile := &Tile{x, y, grid[y][x] == '#', grid[y][x] == '^'}
			row = append(row, tile)

			if grid[y][x] == '^' {
				m.guard = &Guard{x, y, 0, -1, []*Tile{tile}, &m}
				tile.visited = true
			}
		}
		m.tiles = append(m.tiles, row)
	}

	return &m
}
