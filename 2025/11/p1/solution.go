package p1

import (
	"fmt"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

type Graph struct {
	edges map[string][]string
}

func Solve() {
	fmt.Println("--- 11-1 ---")

	strs := helpers.ReadLines("./2025/11/p1/input")
	graph := buildGraph(strs)

	var result int
	result = graph.paths("you")

	fmt.Printf("Result: %d\n", result)
}

func (g *Graph) paths(current string) int {
	paths := 0
	edges := g.edges[current]

	if current != "out" && len(edges) == 0 {
		return 0
	}

	if current == "out" {
		return 1
	}

	for _, edge := range edges {
		paths += g.paths(edge)
	}
	return paths
}

func buildGraph(strs []string) Graph {
	graph := Graph{edges: make(map[string][]string)}
	for _, str := range strs {
		parts := strings.Split(str, ": ")
		edges := strings.Split(parts[1], " ")
		graph.edges[parts[0]] = edges
	}
	return graph
}
