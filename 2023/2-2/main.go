package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []Set
}

type Set map[string]int

func main() {
	fmt.Println("--- 2-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		result += power(game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d", result)
}

func power(game Game) int {
	var maxR, maxG, maxB int = 0, 0, 0
	for _, s := range game.sets {
		if s["red"] > maxR {
			maxR = s["red"]
		}
		if s["green"] > maxG {
			maxG = s["green"]
		}
		if s["blue"] > maxB {
			maxB = s["blue"]
		}
	}

	return maxR * maxG * maxB
}

func parseGame(str string) Game {
	var sets []Set

	split := strings.Split(str, ": ")

	gameStrings := strings.Split(split[0], " ")
	setStrings := strings.Split(split[1], "; ")

	for _, s := range setStrings {
		set := make(map[string]int)
		colors := strings.Split(s, ", ")
		for _, c := range colors {
			split := strings.Split(c, " ")
			set[split[1]], _ = strconv.Atoi(split[0])
		}
		sets = append(sets, set)
	}

	gameId, _ := strconv.Atoi(gameStrings[1])
	return Game{
		id:   gameId,
		sets: sets,
	}
}
