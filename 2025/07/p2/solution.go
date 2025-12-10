package p2

import (
	"fmt"
	"slices"

	"github.com/drcrees/aoc/helpers"
)

var g Graph

type Graph struct {
	vertices map[string][]string
}

func (g *Graph) AddEdge(src, dest string) {
	if g.vertices == nil {
		g.vertices = make(map[string][]string)
	}
	g.vertices[src] = append(g.vertices[src], dest)
}

func Solve() {
	fmt.Println("--- 7-2 ---")

	g = Graph{}
	visited = make(map[string]int)
	grid := helpers.ReadRunes("./2025/07/p2/input")
	result := fireBeam(grid)

	fmt.Printf("Result: %d\n", result)
}

func fireBeam(grid [][]rune) int {
	splits := 0
	for index, r := range grid[0] {
		if string(r) == "S" {
			buildGraph(grid, index, 1, "S")
			break
		}
	}
	splits = countPaths("S") - 1
	return splits
}

var visited map[string]int

func countPaths(vertex string) int {
	count := 0
	if len(g.vertices[vertex]) == 0 {
		visited[vertex] = 2
		return 2
	}
	if len(g.vertices[vertex]) == 1 {
		count++
	}
	for _, v := range g.vertices[vertex] {
		if value, ok := visited[v]; ok {
			count += value
			continue
		}
		visited[v] = count
		count += countPaths(v)
	}
	visited[vertex] = count
	return count
}

func buildGraph(grid [][]rune, x int, y int, current string) {
	for row := y; row < len(grid); row++ {
		if string(grid[row][x]) == "^" {
			if slices.Contains(g.vertices[current], fmt.Sprintf("(%d,%d)", x, row)) {
				break
			}
			g.AddEdge(current, fmt.Sprintf("(%d,%d)", x, row))
			grid[row][x-1] = '|'
			grid[row][x+1] = '|'
			buildGraph(grid, x-1, row+1, fmt.Sprintf("(%d,%d)", x, row))
			buildGraph(grid, x+1, row+1, fmt.Sprintf("(%d,%d)", x, row))
			return
		}
		grid[row][x] = '|'
	}
	return
}
