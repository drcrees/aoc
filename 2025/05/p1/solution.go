package p1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

type Range struct {
	start int
	end   int
}

func Solve() {
	fmt.Println("--- 5-1 ---")

	strs := helpers.ReadLines("./2025/05/p1/input")
	ids := helpers.ReadLines("./2025/05/p1/input2")

	ranges := buildRanges(strs)
	ranges = consolidate(ranges)

	result := countFresh(ranges, parseInts(ids))

	fmt.Printf("Result: %d\n", result)
}

func countFresh(ranges []Range, ids []int) int {
	result := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				result += 1
				break
			}
		}
	}

	return result
}

func consolidate(ranges []Range) []Range {
	var r []Range
	for _, r1 := range ranges {
		skip := false
		for idx, r2 := range r {
			if r1.start >= r2.start && r1.start <= r2.end && r1.end >= r2.end {
				r[idx].start = r2.start
				r[idx].end = r1.end
				skip = true
			}
			if r1.start >= r2.start && r1.end <= r2.end {
				skip = true
			}
		}
		if !skip {
			r = append(r, r1)
		}
	}
	return r
}

func buildRanges(strs []string) []Range {
	var ranges []Range
	for _, r := range strs {
		split := strings.Split(r, "-")
		s, _ := strconv.Atoi(split[0])
		e, _ := strconv.Atoi(split[1])

		ranges = append(ranges, Range{start: s, end: e})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	return ranges
}

func parseInts(strs []string) []int {
	var ids []int
	for _, id := range strs {
		i, _ := strconv.Atoi(id)
		ids = append(ids, i)
	}
	return ids
}
