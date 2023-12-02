package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("--- 1-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int = 0
	for scanner.Scan() {
		result += parse(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d", result)
}

func parse(str string) int {
	var d1 int = -1
	var d2 int = -1

	runes := []rune(str)
	for _, r := range runes {
		if unicode.IsDigit(r) == true {
			if d1 == -1 { // first digit
				d1 = (int(r) - '0')
			}
			d2 = (int(r) - '0')
		}
	}

	result, err := strconv.Atoi(fmt.Sprintf("%d%d", d1, d2))
	if err != nil {
		log.Fatal(err)
	}

	return result
}
