package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var tiles [][]Tile
var pipeline *Pipe

type Tile struct {
	value  string
	isPipe bool
}

type Pipe struct {
	value    string
	distance int
	next     *Pipe

	north *Pipe
	south *Pipe
	east  *Pipe
	west  *Pipe
}

func main() {
	fmt.Println("--- 10-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = 0

	for scanner.Scan() {
		parse(scanner.Text())
	}

	y, x := findS()
	tiles[y][x].isPipe = true

	pipeline = &Pipe{value: "S"}

	pipeline.north = northPipe(x, y, pipeline, 0)
	pipeline.south = southPipe(x, y, pipeline, 0)
	pipeline.east = eastPipe(x, y, pipeline, 0)
	pipeline.west = westPipe(x, y, pipeline, 0)

	result = tilesInsideLoop()
	fmt.Printf("Result: %v", result)
}

func tilesInsideLoop() int {
	var inside = false
	var count = 0

	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[0]); j++ {
			if tiles[i][j].isPipe {
				if tiles[i][j].value == "|" ||
					tiles[i][j].value == "J" ||
					tiles[i][j].value == "L" {
					inside = !inside
				}
			} else if inside {
				count++
			}
		}
	}
	return count
}

func westPipe(x int, y int, current *Pipe, dist int) *Pipe {
	west := max(x-1, 0)
	if tiles[y][west].value == "-" {
		dist++
		tiles[y][west].isPipe = true
		pipe := &Pipe{value: tiles[y][west].value, distance: dist}
		pipe.next = westPipe(west, y, pipe, dist)
		return pipe
	}
	if tiles[y][west].value == "F" {
		dist++
		tiles[y][west].isPipe = true
		pipe := &Pipe{value: tiles[y][west].value, distance: dist}
		pipe.next = southPipe(west, y, pipe, dist)
		return pipe
	}
	if tiles[y][west].value == "L" {
		dist++
		tiles[y][west].isPipe = true
		pipe := &Pipe{value: tiles[y][west].value, distance: dist}
		pipe.next = northPipe(west, y, pipe, dist)
		return pipe
	}
	if tiles[y][west].value == "S" {
		dist++
		tiles[y][west].isPipe = true
		pipe := &Pipe{value: tiles[y][west].value, distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func eastPipe(x int, y int, current *Pipe, dist int) *Pipe {
	east := min(x+1, len(tiles[0]))
	if tiles[y][east].value == "-" {
		dist++
		tiles[y][east].isPipe = true
		pipe := &Pipe{value: tiles[y][east].value, distance: dist}
		pipe.next = eastPipe(east, y, pipe, dist)
		return pipe
	}
	if tiles[y][east].value == "7" {
		dist++
		tiles[y][east].isPipe = true
		pipe := &Pipe{value: tiles[y][east].value, distance: dist}
		pipe.next = southPipe(east, y, pipe, dist)
		return pipe
	}
	if tiles[y][east].value == "J" {
		dist++
		tiles[y][east].isPipe = true
		pipe := &Pipe{value: tiles[y][east].value, distance: dist}
		pipe.next = northPipe(east, y, pipe, dist)
		return pipe
	}
	if tiles[y][east].value == "S" {
		dist++
		tiles[y][east].isPipe = true
		pipe := &Pipe{value: tiles[y][east].value, distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func northPipe(x int, y int, current *Pipe, dist int) *Pipe {
	north := max(y-1, 0)
	if tiles[north][x].value == "7" {
		dist++
		tiles[north][x].isPipe = true
		pipe := &Pipe{value: tiles[north][x].value, distance: dist}
		pipe.next = westPipe(x, north, pipe, dist)
		return pipe
	}
	if tiles[north][x].value == "F" {
		dist++
		tiles[north][x].isPipe = true
		pipe := &Pipe{value: tiles[north][x].value, distance: dist}
		pipe.next = eastPipe(x, north, pipe, dist)
		return pipe
	}
	if tiles[north][x].value == "|" {
		dist++
		tiles[north][x].isPipe = true
		pipe := &Pipe{value: tiles[north][x].value, distance: dist}
		pipe.next = northPipe(x, north, pipe, dist)
		return pipe
	}
	if tiles[north][x].value == "S" {
		dist++
		tiles[north][x].isPipe = true
		pipe := &Pipe{value: tiles[north][x].value, distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func southPipe(x int, y int, current *Pipe, dist int) *Pipe {
	south := min(y+1, len(tiles))
	if tiles[south][x].value == "L" {
		dist++
		tiles[south][x].isPipe = true
		pipe := &Pipe{value: tiles[south][x].value, distance: dist}
		pipe.next = eastPipe(x, south, pipe, dist)
		return pipe
	}
	if tiles[south][x].value == "J" {
		dist++
		tiles[south][x].isPipe = true
		pipe := &Pipe{value: tiles[south][x].value, distance: dist}
		pipe.next = westPipe(x, south, pipe, dist)
		return pipe
	}
	if tiles[south][x].value == "|" {
		dist++
		tiles[south][x].isPipe = true
		pipe := &Pipe{value: tiles[south][x].value, distance: dist}
		pipe.next = southPipe(x, south, pipe, dist)
		return pipe
	}
	if tiles[south][x].value == "S" {
		dist++
		tiles[south][x].isPipe = true
		pipe := &Pipe{value: tiles[south][x].value, distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func findS() (int, int) {
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[0]); j++ {
			if tiles[i][j].value == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func parse(str string) {
	values := strings.Split(str, "")

	var t []Tile
	for _, v := range values {
		t = append(t, Tile{value: v})
	}
	tiles = append(tiles, t)
}
