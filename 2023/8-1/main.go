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
	fmt.Println("--- 8-1 ---")

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

	var result = 0
	var found = false

	current := lookupNode("AAA")

	for !found {
		for _, i := range instructions {
			if i == "L" {
				current = current.left
			}
			if i == "R" {
				current = current.right
			}
			result++

			if current.name == "ZZZ" {
				found = true
				break
			}
		}
	}

	fmt.Printf("Result: %v", result)
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
