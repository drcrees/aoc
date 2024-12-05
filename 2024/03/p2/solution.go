package p2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 3-2 ---")
	str := helpers.ReadFile("./2024/03/p2/input")

	r1 := regexp.MustCompile(`(?s)don't\(\)(.*?)+do\(\)`)
	matches := r1.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		str = strings.Replace(str, match[0], "", 1)
	}

	str = strings.SplitN(str, "don't", 2)[0]

	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	m := r.FindAllStringSubmatch(str, -1)

	result := 0
	for _, match := range m {
		v1, _ := strconv.Atoi(match[1])
		v2, _ := strconv.Atoi(match[2])

		fmt.Printf("mul(%d,%d)\n", v1, v2)
		result += (v1 * v2)
	}

	fmt.Printf("Result: %d\n", result)
}
