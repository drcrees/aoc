package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Coords struct {
	x int
	y int
}

var xmap map[Coords]int

func directions() map[int]Coords {
	return map[int]Coords{
		0: {0, 0}, 1: {1, 1}, 2: {1, -1}, 3: {-1, 1}, 4: {-1, -1},
	}
}

func main() {
	now := time.Now()
	defer func() {
		fmt.Printf("in %s\n", time.Now().Sub(now))
	}()

	fmt.Println("--- 4-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int = 0
	var grid [][]rune
	xmap = make(map[Coords]int)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result = whereXmas(grid)
	fmt.Printf("Result: %d\n", result)
}

func whereXmas(grid [][]rune) (result int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'M' {
				isXmas(grid, x, y, 'A', 0)
			}
		}
	}

	for _, v := range xmap {
		if v >= 2 {
			result++
		}
	}

	return result
}

func inbounds(start Coords, op Coords, bounds Coords) bool {
	if start.x+op.x >= 0 && start.x+op.x <= bounds.x {
		return start.y+op.y >= 0 && start.y+op.y <= bounds.y
	}

	return false
}

func isXmas(grid [][]rune, x int, y int, letter rune, direction int) bool {
	for key, dir := range directions() {
		if inbounds(Coords{x, y}, dir, Coords{len(grid) - 1, len(grid[0]) - 1}) {
			if grid[y+dir.y][x+dir.x] == letter {
				if direction == 0 && isXmas(grid, x+dir.x, y+dir.y, 'S', key) {
					xmap[Coords{x + dir.x, y + dir.y}]++
					continue
				}
				if key == direction {
					if letter == 'S' {
						return true
					}
				}
			}
		}
	}

	return false
}
