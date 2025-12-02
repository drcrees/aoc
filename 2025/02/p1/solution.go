package p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 2-1 ---")
	strs := helpers.ReadDelimitedStrings("./2025/02/p1/input", ",")

	result := 0

	for _, str := range strs[0] {
		result += parseIds(str)
	}

	fmt.Printf("Result: %d\n", result)
}

func parseIds(str string) int {
	ids := strings.Split(str, "-")

	min, _ := strconv.Atoi(ids[0])
	max, _ := strconv.Atoi(ids[1])

	// fmt.Printf("%d-%d\n", min, max)
	return check(min, max, 0)
}

func check(current int, max int, total int) int {
	str := strconv.Itoa(current)

	if len(str)%2 == 0 {
		middle := len(str) / 2
		if str[:middle] == str[middle:] {
			total += current
		}
	}

	if current == max {
		return total
	}

	return check(current+1, max, total)
}
