package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- 9-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = 0

	for scanner.Scan() {
		arr := parse(scanner.Text())
		result += find(arr)
	}

	fmt.Printf("Result: %v", result)
}

func find(arr []int) int {
	var sum = 0
	var diffs []int

	var diffSum = 0
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		diffSum += diff

		if diffSum == 0 && (i+2) == len(arr) {
			return arr[i+1]
		}
		diffs = append(diffs, diff)
	}

	sum += arr[len(arr)-1] + find(diffs)
	return sum
}

func parse(str string) []int {
	var arr []int

	values := strings.Split(str, " ")
	for _, v := range values {
		n, _ := strconv.Atoi(v)
		arr = append(arr, n)
	}

	return arr
}
