package p2

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/drcrees/aoc/helpers"
)

type Page struct {
	id     string
	before map[string]*Page
	after  map[string]*Page
}

type Pages []Page

func (p Pages) Len() int {
	return len(p)
}

func (p Pages) Less(i, j int) bool {
	return p[i].before[p[j].id] != nil
}

func (p Pages) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Solve() {
	fmt.Println("--- 5-2 ---")
	rules := helpers.ReadDelimitedStrings("./2024/05/p2/input1", "|")
	updates := helpers.ReadDelimitedStrings("./2024/05/p2/input2", ",")

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
		var pages Pages
		valid := true
		for i, page := range update {
			for j := i + 1; j < len(update); j++ {
				if ordering[page].after[update[j]] != nil {
					valid = false
					break
				}
			}
			pages = append(pages, *ordering[page])
		}

		if !valid {
			sort.Sort(pages)
			value, _ := strconv.Atoi(pages[int(math.Ceil(float64(len(pages)/2)))].id)
			result += value
		}
	}

	fmt.Printf("Result: %d\n", result)
}
