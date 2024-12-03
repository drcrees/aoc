package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var weird = map[string]string{
	"twone":     "21",
	"eightwo":   "82",
	"eighthree": "83",
	"oneight":   "18",
	"threeight": "38",
	"fiveight":  "58",
	"nineight":  "98",
	"sevenine":  "79",
}

func main() {
	fmt.Println("--- 1-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int = 0
	for scanner.Scan() {
		str := convert(scanner.Text())
		result += parse(str)
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

func convert(str string) string {
	for k, v := range weird {
		str = strings.Replace(str, k, v, -1)
	}

	for k, v := range digits {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}
