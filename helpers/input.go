package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadDelimitedStrings(filePath string, delimiter string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var strs [][]string
	for scanner.Scan() {
		strs = append(strs, strings.Split(scanner.Text(), delimiter))
	}

	return strs
}

func ReadDelimitedInts(filePath string, delimiter string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nums []int
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), delimiter)
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
	}

	return nums
}

func ReadInts(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nums [][]int
	for scanner.Scan() {
		var line []int
		for _, r := range scanner.Text() {
			line = append(line, int(r-'0'))
		}
		nums = append(nums, line)
	}

	return nums
}

func ReadRunes(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var runes [][]rune
	for scanner.Scan() {
		runes = append(runes, []rune(scanner.Text()))
	}

	return runes
}
