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
var almanac Almanac

type Almanac struct {
	maps []AlmanacMap
}

type AlmanacMap struct {
	name     string
	elements []Element
}

type Element struct {
	destination int
	source      int
	length      int
}

func main() {
	fmt.Println("--- 5-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result = 0
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
	for _, seed := range seeds {
		soil := lookup(seed, almanac.maps[0])
		fertilizer := lookup(soil, almanac.maps[1])
		water := lookup(fertilizer, almanac.maps[2])
		light := lookup(water, almanac.maps[3])
		temperature := lookup(light, almanac.maps[4])
		humidity := lookup(temperature, almanac.maps[5])
		location := lookup(humidity, almanac.maps[6])

		locations = append(locations, location)
    
    fmt.Printf("seed: %d, location: %d", seed, location)
    fmt.Println()
	}

	sort.Ints(locations)
	result = locations[0]

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v", result)
}

func lookup(value int, m AlmanacMap) int {
	var dest = -1
	for _, e := range m.elements {
		sourceRange := e.source + e.length

		if value >= e.source && value < sourceRange {
			dest = e.destination + (value - e.source)
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

		elements = append(elements, Element{destination: d, source: s, length: l})
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
}
