package p1

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

type Machine struct {
	id      int
	config  uint
	buttons []uint
}

func Solve() {
	fmt.Println("--- 10-1 ---")

	strs := helpers.ReadLines("./2025/10/p1/input")
	machines := parseMachines(strs)

	var result int
	for _, m := range machines {
		result += fewestPresses(m)
	}

	fmt.Printf("Result: %d\n", result)
}

func fewestPresses(m Machine) int {
	c := combos(m.buttons)

	for _, subset := range c {
		var value uint
		for _, button := range subset {
			value ^= button
		}
		if value == m.config {
			return len(subset)
		}
	}
	return 0
}

func combos(buttons []uint) (subsets [][]uint) {
	length := uint(len(buttons))
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []uint
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, buttons[object])
			}
		}
		subsets = append(subsets, subset)
	}
	sort.Slice(subsets, func(a, b int) bool {
		return len(subsets[a]) < len(subsets[b])
	})
	return subsets
}

func parseMachines(strs []string) []Machine {
	p1 := regexp.MustCompile(`[\.|#]+`)
	p2 := regexp.MustCompile(`\(\S+\)+`)

	var machines []Machine
	for i, str := range strs {
		s1 := p1.FindString(str)
		s2 := p2.FindAllString(str, -1)

		var config uint
		for i := range s1 {
			if s1[i] == '#' {
				config += (1 << i)
			}
		}

		var buttons []uint
		for _, bs := range s2 {
			var button uint
			for _, s := range strings.Split(bs[1:len(bs)-1], ",") {
				v, _ := strconv.Atoi(s)
				button ^= (1 << v)
			}
			buttons = append(buttons, button)
		}
		machines = append(machines, Machine{i + 1, config, buttons})
	}

	return machines
}
