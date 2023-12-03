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
	gear  gear
}

type gear struct {
	x, y int
}

func main() {
	fmt.Println("--- 3-2 ---")

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
	result = calculateGearRatios()
	fmt.Printf("Result: %v", result)
}

func calculateGearRatios() int {
	var result int = 0
	processed := make(map[gear]bool)

	for _, n := range numbers {
		for _, n2 := range numbers {
			v1, _ := strconv.Atoi(string(n.value))
			v2, _ := strconv.Atoi(string(n2.value))

			if !processed[n.gear] && n.gear == n2.gear {
				if v1 != v2 && n.gear.x != 0 && n.gear.y != 0 {
					processed[n.gear] = true
					result += v1 * v2
				}
			}
		}
	}
	return result
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
				number := &number{x: x, y: y, value: engine[y][x:x2]}
				numbers = append(numbers, number)
				if hasSymbols(engine, number, x, x2, y) {
					v, _ := strconv.Atoi(string(engine[y][x:x2]))
					sum += v
				}
				x = x2
			}
		}
	}
	return sum
}

func hasSymbols(engine [][]rune, number *number, startX, endX, rowIndex int) bool {
	above := max(rowIndex-1, 0)
	below := min(rowIndex+2, len(engine))
	left := max(startX-1, 0)
	right := min(endX+1, len(engine[0]))

	for y := above; y < below; y++ {
		for x := left; x < right; x++ {
			r := engine[y][x]
			if !(unicode.IsDigit(r) || string(r) == ".") {
				if string(r) == "*" {
					number.gear = gear{x: x, y: y}
				}
				return true
			}
		}
	}
	return false
}
