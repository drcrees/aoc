package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var engine [][]rune
var numbers []*number

type number struct {
	x, y  int
	value []rune
}

func main() {
	fmt.Println("--- 3-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		engine = append(engine, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := findNumbers(engine)
	fmt.Printf("Result: %v", result)
}

func containsSymbol(chars []rune) bool {
	for _, c := range chars {
		if !unicode.IsDigit(c) && string(c) != "." {
			return true
		}
	}
	return false
}

func findNumbers(engine [][]rune) int {
	var sum = 0
	for y, row := range engine {
		for x := 0; x < len(row); x++ {
			if unicode.IsDigit(row[x]) {
				var x2 = x + 1
				for ; x2 < len(engine[y]); x2++ {
					if !unicode.IsDigit(engine[y][x2]) {
						break
					}
				}
				numbers = append(numbers, &number{x: x, y: y, value: engine[y][x:x2]})
				if hasSymbols(engine, x, x2, y) {
					v, _ := strconv.Atoi(string(engine[y][x:x2]))
					sum += v
				}
				x = x2
			}
		}
	}
	return sum
}

func hasSymbols(engine [][]rune, startX, endX, rowIndex int) bool {
	above := max(rowIndex-1, 0)
	below := min(rowIndex+2, len(engine))
	left := max(startX-1, 0)
	right := min(endX+1, len(engine[0]))

	for y := above; y < below; y++ {
		for x := left; x < right; x++ {
			r := engine[y][x]
			if !(unicode.IsDigit(r) || string(r) == ".") {
				return true
			}
		}
	}
	return false
}
