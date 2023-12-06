package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var seeds []int
var seedRanges []Range

var almanac Almanac

type Almanac struct {
	maps []AlmanacMap
}

type AlmanacMap struct {
	name     string
	elements []Element
}

type Element struct {
	destination Range
	source      Range
}

type Range struct {
	start int
	end   int
}

func main() {
	fmt.Println("--- 5-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parseSeeds(scanner.Text())
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "map:") {
			var name = strings.Replace(scanner.Text(), " map:", "", 1)
			var strs []string
			for scanner.Scan() {
				if scanner.Text() == "" {
					break
				}
				strs = append(strs, scanner.Text())
			}
			almanac.maps = append(almanac.maps, parseMap(name, strs))
		}
	}

	var locations []int

	for _, sr := range seedRanges {
		fmt.Printf("[%d, %d]", sr.start, sr.end)
		fmt.Println()
		for seed := sr.start; seed <= sr.end; seed++ {
			location := lookupLocation(seed)
			locations = append(locations, location)
		}
		sort.Ints(locations)
		fmt.Printf("Minimum location: %d", locations[0])
		fmt.Println()
		locations = make([]int, 0)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func lookupLocation(seed int) int {
	soil := lookup(seed, almanac.maps[0])
	fertilizer := lookup(soil, almanac.maps[1])
	water := lookup(fertilizer, almanac.maps[2])
	light := lookup(water, almanac.maps[3])
	temperature := lookup(light, almanac.maps[4])
	humidity := lookup(temperature, almanac.maps[5])
	return lookup(humidity, almanac.maps[6])
}

func lookup(value int, m AlmanacMap) int {
	var dest = -1
	for _, e := range m.elements {
		if value >= e.source.start && value <= e.source.end {
			dest = e.destination.start + (value - e.source.start)
			break
		}
	}
	if dest == -1 {
		return value
	}
	return dest
}

func parseMap(name string, strs []string) AlmanacMap {
	var elements []Element
	for _, s := range strs {
		values := strings.Split(s, " ")

		d, _ := strconv.Atoi(values[0])
		s, _ := strconv.Atoi(values[1])
		l, _ := strconv.Atoi(values[2])

		elements = append(
			elements,
			Element{
				destination: Range{start: d, end: d + l - 1},
				source:      Range{start: s, end: s + l - 1},
			},
		)
	}
	return AlmanacMap{name: name, elements: elements}
}

func parseSeeds(str string) {
	str = strings.Replace(str, "seeds: ", "", 1)
	seedStrings := strings.Split(str, " ")

	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	for i := 0; i < len(seeds)-1; i += 2 {
		seedRanges = append(
			seedRanges,
			Range{start: seeds[i], end: seeds[i] + seeds[i+1] - 1},
		)
	}
}
