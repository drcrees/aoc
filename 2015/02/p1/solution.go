package p1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 2-1 ---")
	strs := helpers.ReadLines("./2015/02/p1/input")

	result := 0
	for _, str := range strs {
		dimensions := strings.Split(str, "x")

		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		lw := l * w
		wh := w * h
		hl := h * l

		arr := []int{l * w, w * h, h * l}
		sort.Ints(arr)

		result += (2 * lw) + (2 * wh) + (2 * hl) + arr[0]
	}
	fmt.Printf("Result: %d\n", result)
}
