package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var instructions = ""
var nodes []*Node

type Node struct {
	name  string
	left  *Node
	right *Node
}

func main() {
	fmt.Println("--- 8-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := strings.Split(scanner.Text(), "")
	scanner.Scan()

	for scanner.Scan() {
		parseNode(scanner.Text())
	}

	var result = 1

	startingNodes := getStartingNodes()

	var steps []int
	for _, s := range startingNodes {
		steps = append(steps, getSteps(s, instructions))
	}

	for _, s := range steps {
		result = lcm(result, s)
	}

	fmt.Printf("Result: %v", result)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func getSteps(startingNode *Node, instructions []string) int {
	steps := 0
	current := startingNode
	for !found(current) {
		for _, i := range instructions {
			if i == "L" {
				current = current.left
			}
			if i == "R" {
				current = current.right
			}
			steps++
		}
	}

	return steps
}

func found(current *Node) bool {
	if strings.HasSuffix(current.name, "Z") {
		return true
	}

	return false
}

func getStartingNodes() []*Node {
	var startingNodes []*Node
	for _, n := range nodes {
		if strings.HasSuffix(n.name, "A") {
			startingNodes = append(startingNodes, n)
		}
	}
	return startingNodes
}

func lookupNode(name string) *Node {
	for _, n := range nodes {
		if n.name == name {
			return n
		}
	}
	nodes = append(nodes, &Node{name: name})
	return nodes[len(nodes)-1]
}

func parseNode(str string) {
	values := strings.Split(str, " = ")
	name := values[0]

	values[1] = strings.Replace(values[1], "(", "", 1)
	values[1] = strings.Replace(values[1], ")", "", 1)

	nodeNames := strings.Split(values[1], ", ")

	node := lookupNode(name)
	node.left = lookupNode(nodeNames[0])
	node.right = lookupNode(nodeNames[1])
}
