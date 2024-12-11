package p2

import (
	"fmt"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 11-2 ---")

	stones := helpers.ReadDelimitedInts("./2024/11/p2/input", " ")

	result := blinkStones(stones, 75)
	fmt.Printf("Result: %d\n", result)
}

func blinkStones(stones []int, times int) int {
	count := 0
	cache := make(map[[2]int]int)
	for _, stone := range stones {
		count += blinkStone(stone, times, cache)
	}
	return count
}

func blinkStone(stone int, times int, cache map[[2]int]int) int {
	if c, ok := cache[[2]int{stone, times}]; ok {
		return c
	}

	next := blink(stone)
	cache[[2]int{stone, 1}] = len(next)
	if times == 1 {
		return len(next)
	}

	count := 0
	for _, v := range next {
		count += blinkStone(v, times-1, cache)
	}
	cache[[2]int{stone, times}] = count
	return count
}

func blink(stone int) []int {
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
