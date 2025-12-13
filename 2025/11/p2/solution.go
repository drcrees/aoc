package p2

import (
	"fmt"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

type Graph struct {
	edges map[string][]string
}

func Solve() {
	fmt.Println("--- 11-2 ---")

	visited = make(map[node]int)
	paths = [][]string{}

	strs := helpers.ReadLines("./2025/11/p2/input")
	graph := buildGraph(strs)

	var result int
	result = graph.paths("svr", "out", false, false)

	fmt.Printf("Result: %d\n", result)
}

type node struct {
	name string
	dac  bool
	fft  bool
}

var visited map[node]int
var paths [][]string

func (g *Graph) paths(current string, destination string, dac bool, fft bool) int {
	count := 0
	edges := g.edges[current]

	node := node{current, dac, fft}

	if current == destination {
		if dac && fft {
			return 1
		}
		return 0
	}

	if res, ok := visited[node]; ok {
		return res
	}

	for _, edge := range edges {
		count += g.paths(edge, destination, dac || edge == "dac", fft || edge == "fft")
	}
	visited[node] = count
	return count
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

func (g *Graph) print() {
	for node, edge := range g.edges {
		fmt.Printf("%s -> %s\n", node, edge)
	}
}
