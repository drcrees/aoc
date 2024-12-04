package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Letter struct {
	x         int
	y         int
	value     rune
	direction int
}

func main() {
	now := time.Now()
	defer func() {
		fmt.Printf("in %s\n", time.Now().Sub(now))
	}()

	fmt.Println("--- 4-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int = 0
	var grid [][]rune
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
			if grid[y][x] == 'X' {
				count++
				result += isXmas(grid, x, y, 'M', 0)
			}
		}
	}

	return result
}

func isXmas(grid [][]rune, x int, y int, letter rune, direction int) int {
	numRows := len(grid) - 1
	numCols := len(grid[0]) - 1

	numXmas := 0

	if x-1 >= 0 && grid[y][x-1] == letter {
		if direction == 0 {
			numXmas += isXmas(grid, x-1, y, 'A', 1)
		} else if direction == 1 {
			if letter == 'S' {
				return 1
			}
			numXmas += isXmas(grid, x-1, y, 'S', 1)
		}
	}

	if x+1 <= numCols && grid[y][x+1] == letter {
		if direction == 0 {
			numXmas += isXmas(grid, x+1, y, 'A', 2)
		} else if direction == 2 {
			if letter == 'S' {
				return 1
			}
			numXmas += isXmas(grid, x+1, y, 'S', 2)
		}
	}

	if y-1 >= 0 {
		if grid[y-1][x] == letter {
			if direction == 0 {
				numXmas += isXmas(grid, x, y-1, 'A', 3)
			} else if direction == 3 {
				if letter == 'S' {
					return 1
				}
				numXmas += isXmas(grid, x, y-1, 'S', 3)
			}
		}
		if x-1 >= 0 && grid[y-1][x-1] == letter {
			if direction == 0 {
				numXmas += isXmas(grid, x-1, y-1, 'A', 4)
			} else if direction == 4 {
				if letter == 'S' {
					return 1
				}
				numXmas += isXmas(grid, x-1, y-1, 'S', 4)
			}
		}
		if x+1 <= numCols && grid[y-1][x+1] == letter {
			if direction == 0 {
				numXmas += isXmas(grid, x+1, y-1, 'A', 5)
			} else if direction == 5 {
				if letter == 'S' {
					return 1
				}
				numXmas += isXmas(grid, x+1, y-1, 'S', 5)
			}
		}
	}

	if y+1 <= numRows {
		if grid[y+1][x] == letter {
			if direction == 0 {
				numXmas += isXmas(grid, x, y+1, 'A', 6)
			} else if direction == 6 {
				if letter == 'S' {
					return 1
				}
				numXmas += isXmas(grid, x, y+1, 'S', 6)
			}
		}
		if x-1 >= 0 && grid[y+1][x-1] == letter {
			if direction == 0 {
				numXmas += isXmas(grid, x-1, y+1, 'A', 7)
			} else if direction == 7 {
				if letter == 'S' {
					return 1
				}
				numXmas += isXmas(grid, x-1, y+1, 'S', 7)
			}
		}
		if x+1 <= numCols && grid[y+1][x+1] == letter {
			if direction == 0 {
				numXmas += isXmas(grid, x+1, y+1, 'A', 8)
			} else if direction == 8 {
				if letter == 'S' {
					return 1
				}
				numXmas += isXmas(grid, x+1, y+1, 'S', 8)
			}
		}
	}

	return numXmas
}
