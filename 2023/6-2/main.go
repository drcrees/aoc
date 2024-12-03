package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var races []Race

type Race struct {
	time     int
	distance int
}

func main() {
	fmt.Println("--- 6-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	t := scanner.Text()
	scanner.Scan()
	d := scanner.Text()

	var result = 1

	parse(t, d)
	for _, r := range races {
		result *= calcPotential(r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v", result)
}

func calcPotential(r Race) int {
	var potential = 0
	for i := 0; i <= r.time; i++ {
		remaining := r.time - i
		if (i * remaining) > r.distance {
			potential++
		}
	}
	return potential
}

func parse(timeStr string, distanceStr string) {
	times := strings.Split(timeStr, " ")
	distances := strings.Split(distanceStr, " ")

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])

		races = append(races, Race{time: t, distance: d})
	}
}
