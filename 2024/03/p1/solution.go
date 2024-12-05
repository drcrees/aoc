package p1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 3-1 ---")
	str := helpers.ReadFile("./2024/03/p1/input")

	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	matches := r.FindAllStringSubmatch(str, -1)

	result := 0
	for _, match := range matches {
		v1, _ := strconv.Atoi(match[1])
		v2, _ := strconv.Atoi(match[2])

		result += (v1 * v2)
	}

	fmt.Printf("Result: %d\n", result)
}
