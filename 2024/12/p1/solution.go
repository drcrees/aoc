package p1

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

// var all_directions []Coords = []Coords{{1,0},{-1,0},{0,1},{0,-1},{1,1},{1,-1},{-1,1},{-1,-1}}
var directions []Coords = []Coords{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

type Coords struct {
	x, y int
}

type Edge struct {
	coords    Coords
	direction Coords
}

type Region struct {
	id    string
	plots []*Plot
	sides int
}

type Plot struct {
	plant     rune
	coords    Coords
	region    *Region
	visited   bool
	perimeter int
	edges     []Edge
}

func Solve() {
	fmt.Println("--- 12-1 ---")

	grid := helpers.ReadRunes("./2024/12/p1/input")

	plots := buildPlots(grid)
	regions := buildRegions(plots)

	result := 0
	for _, region := range regions {
		perimeter := 0

		for _, plot := range region.plots {
			perimeter += plot.perimeter
		}

		region.sides = sides(region.plots)

		// result += len(region.plots) * perimeter
		result += region.sides * len(region.plots)
		fmt.Printf("%s - area: %d, perimeter: %d, sides: %d\n", region.id, len(region.plots), perimeter, region.sides)
	}

	fmt.Printf("Result: %d\n", result)
}

func sides(plots []*Plot) int {
	var all []Edge

	currentDirection := 0
	for _, p := range plots {
		for _, e := range p.edges {
			if e.direction.x == directions[currentDirection].x && e.direction.y == directions[currentDirection].y {
				all = append(all, p.edges...)
				if currentDirection == 3 {
					currentDirection = 0
				} else {
					currentDirection++
				}
			} else {
				currentDirection--
			}
		}
	}

	fmt.Println(all)
	count := 0
	for i := 0; i < len(all)-1; i++ {
		fmt.Printf("(%d, %d)\n", all[i].direction.x, all[i].direction.y)
		if all[i].direction.x != all[i+1].direction.x || all[i].direction.y != all[i+1].direction.y {
			count++
		}
	}

	return 0
}

func buildPlots(grid [][]rune) [][]*Plot {
	plots := make([][]*Plot, len(grid))

	for y := 0; y < len(grid); y++ {
		plots[y] = make([]*Plot, len(grid[0]))

		for x := 0; x < len(grid[0]); x++ {
			plots[y][x] = &Plot{
				plant:   grid[y][x],
				coords:  Coords{x, y},
				region:  nil,
				visited: false,
			}
		}
	}

	return plots
}

func buildRegions(plots [][]*Plot) []*Region {
	var regions []*Region
	for y := 0; y < len(plots); y++ {
		for x := 0; x < len(plots[0]); x++ {
			if !plots[y][x].visited {
				regions = append(regions, buildRegion(plots[y][x], plots))
			}
		}
	}

	return regions
}

func buildRegion(plot *Plot, plots [][]*Plot) *Region {
	var p []*Plot

	id := strconv.Itoa(plot.coords.x) + strconv.Itoa(plot.coords.y) + "-" + string(plot.plant)
	p = append(p, findNeighbors(plot, plots)...)

	return &Region{
		id:    id,
		plots: p,
	}
}

func findNeighbors(plot *Plot, plots [][]*Plot) []*Plot {
	var p []*Plot

	p = append(p, plot)
	plot.visited = true

	for _, delta := range directions {
		dx := plot.coords.x + delta.x
		dy := plot.coords.y + delta.y
		if inbounds(plot.coords, delta, plots) && plots[dy][dx].plant == plot.plant {
			if !plots[dy][dx].visited {
				neighbors := findNeighbors(plots[dy][dx], plots)
				p = append(p, neighbors...)
			}
			continue
		}

		if !inbounds(plot.coords, delta, plots) || plots[dy][dx].plant != plot.plant {
			plot.perimeter++
			plot.edges = append(plot.edges, Edge{Coords{plot.coords.x, plot.coords.y}, delta})
		}
	}

	return p
}

func getVisited(plots []*Plot) int {
	num := 0
	for _, v := range plots {
		if v.visited {
			num++
		}
	}

	return num
}

func sameDirection(current Coords, delta Coords) bool {
	if current != delta {
		return false
	}
	return true
}

func inbounds(current Coords, delta Coords, plots [][]*Plot) bool {
	if current.x+delta.x >= 0 && current.x+delta.x < len(plots[0]) {
		return current.y+delta.y >= 0 && current.y+delta.y < len(plots)
	}

	return false
}
