package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var image [][]Pixel
var galaxies []Galaxy

type Pixel struct {
	value    string
	isGalaxy bool
}

type Galaxy struct {
	x int
	y int
}

func main() {
	fmt.Println("--- 11-1 ---")

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

	expandColumns()
	expandRows()
	findGalaxies()

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += shortestDistance(galaxies[i], galaxies[j])
		}
	}

	fmt.Printf("Result: %v", result)
}

func printRow(row int) {
	for i := 0; i < len(image[0]); i++ {
		fmt.Printf("%s", image[row][i].value)
	}
	fmt.Println()
}

func shortestDistance(i Galaxy, j Galaxy) int {
	var distance = 0

	maxX := max(i.x, j.x)
	minX := min(i.x, j.x)
	maxY := max(i.y, j.y)
	minY := min(i.y, j.y)

	distance += (maxX - minX) * 2
	distance += (maxY - minY) - (maxX - minX)

	return distance
}

func findGalaxies() {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j].value == "#" {
				galaxies = append(galaxies, Galaxy{x: j, y: i})
			}
		}
	}
}

func expandRows() {
	var indices []int
	for i := 0; i < len(image); i++ {
		if image[i][0].value == "." {
			var expand = true
			for j := 0; j < len(image[0]); j++ {
				if image[i][j].value == "#" {
					expand = false
				}
			}
			if expand {
				indices = append(indices, i)
			}
		}
	}

	for i, index := range indices {
		var newRow []Pixel
		index = i + index
		for i := 0; i < len(image[0]); i++ {
			newRow = append(newRow, Pixel{value: "."})
		}
		image = append(image[:index+1], image[index:]...)
		image[index] = newRow
	}
}

func expandColumns() {
	var indices []int
	for i := 0; i < len(image[0]); i++ {
		if image[0][i].value == "." {
			var expand = true
			for j := 0; j < len(image); j++ {
				if image[j][i].value == "#" {
					expand = false
				}
			}
			if expand {
				indices = append(indices, i)
			}
		}
	}

	for i, index := range indices {
		index = i + index
		for j := 0; j < len(image); j++ {
			image[j] = append(image[j][:index+1], image[j][index:]...)
			image[j][index] = Pixel{value: "."}
		}
	}
	fmt.Println()
}

func parse(str string) {
	pixelValues := strings.Split(str, "")

	var pixels []Pixel
	for _, p := range pixelValues {
		pixels = append(pixels, Pixel{value: p, isGalaxy: p == "#"})
	}
	image = append(image, pixels)
}
