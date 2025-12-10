package p2

import (
	"cmp"
	"fmt"
	"github.com/drcrees/aoc/helpers"
	"maps"
	"slices"
	"strconv"
)

type Box struct {
	x, y, z int
}

type Edge struct {
	from     Box
	to       Box
	distance int
}

func distance(b1, b2 Box) int {
	return (b2.x-b1.x)*(b2.x-b1.x) + (b2.y-b1.y)*(b2.y-b1.y) + (b2.z-b1.z)*(b2.z-b1.z)
}

func Solve() {
	fmt.Println("--- 8-2 ---")
	result := 0

	strs := helpers.ReadDelimitedStrings("./2025/08/p2/input", ",")

	boxes := boxes(strs)
	edges := distances(boxes)
	circuits := []map[Box]bool{}

	for _, box := range boxes {
		circuits = append(circuits, map[Box]bool{box: true})
	}

	for i, e := range edges {

		var c1, c2 int
		for i, c := range circuits {
			if c[e.from] == true {
				c1 = i
			}
			if c[e.to] == true {
				c2 = i
			}
		}

		if c1 != c2 {
			maps.Copy(circuits[c1], circuits[c2])
			circuits = slices.Delete(circuits, c2, c2+1)
		}

		if i+1 == 1000 {
			slices.SortFunc(circuits, func(a, b map[Box]bool) int {
				return -cmp.Compare(len(a), len(b))
			})
		}

		if len(circuits) == 1 {
			result = (e.from.x * e.to.x)
			break
		}
	}
	fmt.Printf("Result: %d\n", result)
}

var lookup map[Box][]Box

func distances(boxes []Box) (edges []Edge) {
	lookup := make(map[Box][]Box)
	for _, b1 := range boxes {
		for _, b2 := range boxes {
			distance := distance(b1, b2)
			if distance != 0 {
				found := false
				if exists, ok := lookup[b2]; ok {
					for _, b := range exists {
						if b == b1 {
							found = true
							break
						}
					}
				}
				if !found {
					lookup[b1] = append(lookup[b1], b2)
					edges = append(edges, Edge{b1, b2, distance})
				}
			}
		}
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return edges
}

func boxes(strs [][]string) []Box {
	var boxes []Box
	for _, str := range strs {
		var x, y, z int
		x, _ = strconv.Atoi(str[0])
		y, _ = strconv.Atoi(str[1])
		z, _ = strconv.Atoi(str[2])
		boxes = append(boxes, Box{x, y, z})
	}
	return boxes
}
