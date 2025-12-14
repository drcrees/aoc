package p1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape struct {
	layout []string
}

type Region struct {
	grid   [][]string
	counts []int
}

func Solve() {
	fmt.Println("--- 12-1 ---")

	shapes, regions := parse()

	var result int
	for _, r := range regions {
		var areaCheck, boundsCheck bool
		minSpaceRequired := 0
		maxSpaceRequired := 0
		for i, numShapes := range r.counts {
			required := shapes[i].area() * numShapes
			minSpaceRequired += required
		}
		if minSpaceRequired <= r.area() {
			areaCheck = true
		}
		for i, numShapes := range r.counts {
			required := shapes[i].bounds() * numShapes
			maxSpaceRequired += required
		}
		if maxSpaceRequired <= r.area() {
			boundsCheck = true
		}
		if areaCheck && boundsCheck {
			result += 1
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func (s *Shape) bounds() int {
	return len(s.layout) * len(s.layout[0])
}

func (s *Shape) area() int {
	area := 0
	for _, row := range s.layout {
		for _, c := range row {
			if c == '#' {
				area++
			}
		}
	}
	return area
}

func (r *Region) area() int {
	area := 0
	for _, row := range r.grid {
		area += len(row)
	}
	return area
}

func parse() ([]Shape, []Region) {
	var shapes []Shape
	var regions []Region

	file, err := os.Open("./2025/12/p1/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var shape Shape
	var region Region
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '.' || line[0] == '#' {
			shape.layout = append(shape.layout, line)
			for scanner.Scan() {
				if strings.Trim(scanner.Text(), "") == "" {
					shapes = append(shapes, shape)
					shape = Shape{}
					break
				} else {
					shape.layout = append(shape.layout, scanner.Text())
				}
			}
		}
		if len(strings.Split(line, ": ")) > 1 {
			parts := strings.Split(line, ": ")
			sizes := strings.Split(parts[0], "x")
			indices := strings.Split(parts[1], " ")

			region = Region{}
			x, _ := strconv.Atoi(sizes[0])
			y, _ := strconv.Atoi(sizes[1])
			r := []string{}
			for row := 0; row < y; row++ {
				for col := 0; col < x; col++ {
					r = append(r, ".")
				}
				region.grid = append(region.grid, r)
				r = []string{}
			}
			for _, i := range indices {
				index, _ := strconv.Atoi(i)
				region.counts = append(region.counts, index)
			}
			regions = append(regions, region)
		}
		lines = append(lines, scanner.Text())
	}

	return shapes, regions
}
