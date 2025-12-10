package p1

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
	fmt.Println("--- 7-1 ---")

	g = Graph{}
	grid := helpers.ReadRunes("./2025/07/p1/input")
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
	splits = countSplits("S", make(map[string]bool))
	return splits
}

func countSplits(vertex string, visited map[string]bool) int {
	count := 0
	for _, v := range g.vertices[vertex] {
		if !visited[v] {
			count++
			visited[v] = true
			count += countSplits(v, visited)
		}
	}
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
