package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 2-2 ---")
	strs := helpers.ReadDelimitedStrings("./2025/02/p2/input", ",")

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

	return check(min, max, 0)
}

func check(current int, max int, total int) int {
	str := strconv.Itoa(current)

	strstr := str + str
	s := strstr[1 : len(strstr)-1]

	if strings.Contains(s, str) {
		total += current
	}

	if current == max {
		return total
	}

	return check(current+1, max, total)
}
