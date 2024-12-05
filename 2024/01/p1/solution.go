package p1

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 1-1 ---")
	strs := helpers.ReadDelimitedStrings("./2024/01/p1/input", "   ")

	var l1, l2 []int
	for _, str := range strs {
		n1, _ := strconv.Atoi(str[0])
		n2, _ := strconv.Atoi(str[1])
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	result := 0
	for i := 0; i < len(l1); i++ {
		d := math.Abs(float64(l1[i] - l2[i]))
		result += int(d)
	}

	fmt.Printf("Result: %d\n", result)
}
