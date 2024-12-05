package p1

import (
	"fmt"
	"math"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

type Page struct {
	id     string
	before map[string]*Page
	after  map[string]*Page
}

func Solve() {
	fmt.Println("--- 5-1 ---")
	rules := helpers.ReadDelimitedStrings("./2024/05/p1/input1", "|")
	updates := helpers.ReadDelimitedStrings("./2024/05/p1/input2", ",")

	ordering := make(map[string]*Page)
	for _, rule := range rules {
		r1 := ordering[rule[0]]
		r2 := ordering[rule[1]]

		if r1 == nil {
			r1 = &Page{
				id:     rule[0],
				before: make(map[string]*Page),
				after:  make(map[string]*Page),
			}
			ordering[rule[0]] = r1
		}

		if r2 == nil {
			r2 = &Page{
				id:     rule[1],
				before: make(map[string]*Page),
				after:  make(map[string]*Page),
			}
			ordering[rule[1]] = r2
		}

		r1.before[rule[1]] = ordering[rule[1]]
		r2.after[rule[0]] = ordering[rule[0]]
	}

	result := 0
	for _, update := range updates {
		valid := true
		for i, page := range update {
			for j := i + 1; j < len(update); j++ {
				if ordering[page].after[update[j]] != nil {
					valid = false
				}
			}
		}

		if valid {
			value, _ := strconv.Atoi(update[int(math.Ceil(float64(len(update)/2)))])
			result += value
		}
	}

	fmt.Printf("Result: %d\n", result)
}
