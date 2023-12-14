package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var tiles [][]string
var pipeline *Pipe

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
	fmt.Println("--- 10-1 ---")

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

	pipeline = &Pipe{value: "S"}

	pipeline.north = northPipe(x, y, pipeline, 0)
	pipeline.south = southPipe(x, y, pipeline, 0)
	pipeline.east = eastPipe(x, y, pipeline, 0)
	pipeline.west = westPipe(x, y, pipeline, 0)

	n := findMaxDistance(pipeline.north)
	s := findMaxDistance(pipeline.south)
	e := findMaxDistance(pipeline.east)
	w := findMaxDistance(pipeline.west)

	result = max(n, s, e, w) / 2

	fmt.Printf("Result: %v", result)
}

func findMaxDistance(pipeline *Pipe) int {
	var max = 0
	for pipeline != nil && pipeline.next != nil {
		if pipeline.next.distance > max {
			max = pipeline.next.distance
		}
		pipeline = pipeline.next
	}
	return max
}

func westPipe(x int, y int, current *Pipe, dist int) *Pipe {
	west := max(x-1, 0)
	if tiles[y][west] == "-" {
		dist++
		pipe := &Pipe{value: tiles[y][west], distance: dist}
		pipe.next = westPipe(west, y, pipe, dist)
		return pipe
	}
	if tiles[y][west] == "F" {
		dist++
		pipe := &Pipe{value: tiles[y][west], distance: dist}
		pipe.next = southPipe(west, y, pipe, dist)
		return pipe
	}
	if tiles[y][west] == "L" {
		dist++
		pipe := &Pipe{value: tiles[y][west], distance: dist}
		pipe.next = northPipe(west, y, pipe, dist)
		return pipe
	}
	if tiles[y][west] == "S" {
		dist++
		pipe := &Pipe{value: tiles[y][west], distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func eastPipe(x int, y int, current *Pipe, dist int) *Pipe {
	east := min(x+1, len(tiles[0]))
	if tiles[y][east] == "-" {
		dist++
		pipe := &Pipe{value: tiles[y][east], distance: dist}
		pipe.next = eastPipe(east, y, pipe, dist)
		return pipe
	}
	if tiles[y][east] == "7" {
		dist++
		pipe := &Pipe{value: tiles[y][east], distance: dist}
		pipe.next = southPipe(east, y, pipe, dist)
		return pipe
	}
	if tiles[y][east] == "J" {
		dist++
		pipe := &Pipe{value: tiles[y][east], distance: dist}
		pipe.next = northPipe(east, y, pipe, dist)
		return pipe
	}
	if tiles[y][east] == "S" {
		dist++
		pipe := &Pipe{value: tiles[y][east], distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func northPipe(x int, y int, current *Pipe, dist int) *Pipe {
	north := max(y-1, 0)
	if tiles[north][x] == "7" {
		dist++
		pipe := &Pipe{value: tiles[north][x], distance: dist}
		pipe.next = westPipe(x, north, pipe, dist)
		return pipe
	}
	if tiles[north][x] == "F" {
		dist++
		pipe := &Pipe{value: tiles[north][x], distance: dist}
		pipe.next = eastPipe(x, north, pipe, dist)
		return pipe
	}
	if tiles[north][x] == "|" {
		dist++
		pipe := &Pipe{value: tiles[north][x], distance: dist}
		pipe.next = northPipe(x, north, pipe, dist)
		return pipe
	}
	if tiles[north][x] == "S" {
		dist++
		pipe := &Pipe{value: tiles[north][x], distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func southPipe(x int, y int, current *Pipe, dist int) *Pipe {
	south := min(y+1, len(tiles))
	if tiles[south][x] == "L" {
		dist++
		pipe := &Pipe{value: tiles[south][x], distance: dist}
		pipe.next = eastPipe(x, south, pipe, dist)
		return pipe
	}
	if tiles[south][x] == "J" {
		dist++
		pipe := &Pipe{value: tiles[south][x], distance: dist}
		pipe.next = westPipe(x, south, pipe, dist)
		return pipe
	}
	if tiles[south][x] == "|" {
		dist++
		pipe := &Pipe{value: tiles[south][x], distance: dist}
		pipe.next = southPipe(x, south, pipe, dist)
		return pipe
	}
	if tiles[south][x] == "S" {
		dist++
		pipe := &Pipe{value: tiles[south][x], distance: dist}
		pipe.next = pipeline
		return pipe
	}
	return nil
}

func findS() (int, int) {
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[0]); j++ {
			if tiles[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func parse(str string) {
	values := strings.Split(str, "")
	tiles = append(tiles, values)
}
