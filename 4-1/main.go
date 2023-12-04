package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []string
	numbers        map[string]int
}

func main() {
	fmt.Println("--- 4-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := parse(scanner.Text())
		result += calculateWinnings(c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v", result)
}

func calculateWinnings(c Card) int {
	var winnings = 0
	for _, n := range c.winningNumbers {
		if c.numbers[n] > 0 {
			if winnings == 0 { // first match
				winnings += 1
				continue
			}
			winnings = winnings * 2
		}
	}
	return winnings
}

func parse(str string) Card {
	str = strings.Replace(str, "  ", " ", -1) // sanitize
	split := strings.Split(str, ": ")

	numberStrings := strings.Split(split[1], " | ")

	winningNumbers := strings.Split(numberStrings[0], " ")
	numbers := strings.Split(numberStrings[1], " ")

	m := make(map[string]int)
	for _, n := range numbers {
		m[n], _ = strconv.Atoi(n)
	}

	return Card{
		winningNumbers: winningNumbers,
		numbers:        m,
	}
}
