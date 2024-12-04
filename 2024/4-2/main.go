package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Coord struct {
	x int
	y int
}

var xmap map[Coord]int

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
	xmap = make(map[Coord]int)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result = whereXmas(grid)
	fmt.Printf("Result: %d\n", result)
}

func whereXmas(grid [][]rune) int {
	numRows := len(grid)
	numCols := len(grid[0])

	result := 0
	count := 0
	for y := 0; y < numRows; y++ {
		for x := 0; x < numCols; x++ {
			if grid[y][x] == 'M' {
				count++
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

func isXmas(grid [][]rune, x int, y int, letter rune, direction int) bool {
	numRows := len(grid) - 1
	numCols := len(grid[0]) - 1

	if y-1 >= 0 {
		if x-1 >= 0 && grid[y-1][x-1] == letter {
			if direction == 0 {
				if isXmas(grid, x-1, y-1, 'S', 4) {
					xmap[Coord{x - 1, y - 1}]++
				}
			} else if direction == 4 {
				if letter == 'S' {
					return true
				}
			}
		}
		if x+1 <= numCols && grid[y-1][x+1] == letter {
			if direction == 0 {
				if isXmas(grid, x+1, y-1, 'S', 5) {
					xmap[Coord{x + 1, y - 1}]++
				}
			} else if direction == 5 {
				if letter == 'S' {
					return true
				}
			}
		}
	}

	if y+1 <= numRows {
		if x-1 >= 0 && grid[y+1][x-1] == letter {
			if direction == 0 {
				if isXmas(grid, x-1, y+1, 'S', 7) {
					xmap[Coord{x - 1, y + 1}]++
				}
			} else if direction == 7 {
				if letter == 'S' {
					return true
				}
			}
		}
		if x+1 <= numCols && grid[y+1][x+1] == letter {
			if direction == 0 {
				if isXmas(grid, x+1, y+1, 'S', 8) {
					xmap[Coord{x + 1, y + 1}]++
				}
			} else if direction == 8 {
				if letter == 'S' {
					return true
				}
			}
		}
	}

	return false
}
