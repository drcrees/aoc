package p2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 2-2 ---")
	strs := helpers.ReadLines("./2015/02/p2/input")

	result := 0
	for _, str := range strs {
		dimensions := strings.Split(str, "x")

		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		arr := []int{l * 2, w * 2, h * 2}
		sort.Ints(arr)

		result += (l * w * h) + arr[0] + arr[1]
	}
	fmt.Printf("Result: %d\n", result)
}
