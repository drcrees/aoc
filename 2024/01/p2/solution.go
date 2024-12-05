package p2

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

func Solve() {
	fmt.Println("--- 1-2 ---")
	strs := helpers.ReadDelimitedStrings("./2024/01/p2/input", "   ")

	var m map[int]int = make(map[int]int)

	var l1, l2 []int
	for _, str := range strs {
		n1, _ := strconv.Atoi(str[0])
		n2, _ := strconv.Atoi(str[1])
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	for i := 0; i < len(l1); i++ {
		n := l1[i]

		count := 0
		for j := 0; j < len(l2); j++ {
			if n == l2[j] {
				count++
			}
		}
		m[n] = count
	}

	result := 0
	for i, v := range m {
		result += i * v
	}

	fmt.Printf("Result: %d\n", result)
}
