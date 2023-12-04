package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cardMap map[int]Card

type Card struct {
	id             int
	winningNumbers []string
	numbers        map[string]int
}

func main() {
	fmt.Println("--- 4-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result = 0
	cardMap = make(map[int]Card)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := parse(scanner.Text())
		cardMap[c.id] = c
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, c := range cardMap {
		result += 1
		result += calculateWinnings(c)
	}

	fmt.Printf("Result: %v", result)
}

func calculateWinnings(c Card) int {
	var numMatches = 0
	var runningTotal = 0
	for _, n := range c.winningNumbers {
		if c.numbers[n] > 0 { // match
			numMatches++
			runningTotal++
			runningTotal += calculateWinnings(cardMap[c.id+numMatches])
		}
	}

	return runningTotal
}

func parse(str string) Card {
	str = strings.Replace(str, "   ", " ", -1) // sanitize
	str = strings.Replace(str, "  ", " ", -1)
	split := strings.Split(str, ": ")

	cardStrings := strings.Split(split[0], " ")
	numberStrings := strings.Split(split[1], " | ")

	winningNums := strings.Split(numberStrings[0], " ")
	numbers := strings.Split(numberStrings[1], " ")

	m := make(map[string]int)
	for _, n := range numbers {
		m[n], _ = strconv.Atoi(n)
	}

	cardId, _ := strconv.Atoi(cardStrings[1])

	return Card{
		id:             cardId,
		winningNumbers: winningNums,
		numbers:        m,
	}
}
