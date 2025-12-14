package p2

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"

	"github.com/draffensperger/golp"
)

type Machine struct {
	id       int
	config   uint
	buttons  []Button
	joltages []uint
}

type Button struct {
	indices []uint
}

func Solve() {
	fmt.Println("--- 10-2 ---")

	strs := helpers.ReadLines("./2025/10/p2/input")
	machines := parseMachines(strs)

	var result int
	for _, m := range machines {
		lp := golp.NewLP(0, len(m.buttons))

		objective := make([]float64, len(m.buttons))

		for i := 0; i < len(m.buttons); i++ {
			objective[i] = 1.0
			lp.SetInt(i, true)
		}

		lp.SetObjFn(objective)

		for i := 0; i < len(m.joltages); i++ {
			var entries []golp.Entry
			for j, b := range m.buttons {
				if slices.Contains(b.indices, uint(i)) {
					entries = append(entries, golp.Entry{Col: j, Val: 1.0})
				}
			}
			lp.AddConstraintSparse(entries, golp.EQ, float64(m.joltages[i]))
		}

		lp.Solve()

		for _, v := range lp.Variables() {
			result += int(v)
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func parseMachines(strs []string) []Machine {
	p1 := regexp.MustCompile(`[\.|#]+`)
	p2 := regexp.MustCompile(`\(\S+\)+`)
	p3 := regexp.MustCompile(`{(\S+)}`)

	var machines []Machine
	for i, str := range strs {
		s1 := p1.FindString(str)
		s2 := p2.FindAllString(str, -1)
		s3 := p3.FindString(str)

		var config uint
		for i := range s1 {
			if s1[i] == '#' {
				config += (1 << i)
			}
		}

		var buttons []Button
		for _, bs := range s2 {
			var indices []uint
			for _, s := range strings.Split(bs[1:len(bs)-1], ",") {
				v, _ := strconv.Atoi(s)
				indices = append(indices, uint(v))
			}
			buttons = append(buttons, Button{indices})
		}

		var joltages []uint
		for _, s := range strings.Split(s3[1:len(s3)-1], ",") {
			v, _ := strconv.Atoi(s)
			joltages = append(joltages, uint(v))
		}
		machines = append(machines, Machine{i + 1, config, buttons, joltages})
	}

	return machines
}
