package p1

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 11-1 ---")

	stones := helpers.ReadDelimitedInts("./2024/11/p1/input", " ")

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	result := len(stones)
	fmt.Printf("Result: %d\n", result)
}

func blink(stones []int) []int {
	var newStones []int
	for _, stone := range stones {
		newStones = append(newStones, evolve(stone)...)
	}
	return newStones
}

func evolve(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		n1, _ := strconv.Atoi(s[:len(s)/2])
		n2, _ := strconv.Atoi(s[len(s)/2:])
		return []int{n1, n2}
	}

	return []int{stone * 2024}
}
