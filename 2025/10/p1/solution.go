package p1

import (
	"fmt"
  "regexp"
  "strings"
  "strconv"
  "slices"

	"github.com/drcrees/aoc/helpers"
)

type Machine struct {
  config uint
  buttons []uint
}

func Solve() {
	fmt.Println("--- 10-1 ---")

	strs := helpers.ReadLines("./2025/10/p1/input")
  machines := parseMachines(strs)

	result := 0
  for _, m := range machines {
    result += fewestPresses(m)
  }

  // .##.
  // 0123
  // desired += (1 << index) where index == '#'

  // (0,2) (0,1) = 6
  // (1 << 0 , 1 << 2), (1 << 0, 1 << 1)
  // (1 << 0) ^ (1 << 2) ^ (1 << 0) ^ (1 << 1) = 6

  // fmt.Println((1 << 0) ^ (1 << 2))
  // fmt.Println((1 << 0) ^ (1 << 1))


	fmt.Printf("Result: %d\n", result)
}

func fewestPresses(m Machine) int {
  if slices.Contains(m.buttons, m.config) {
    return 1
  }
  return search(m.config, m.buttons, 2)
}

func search(config uint, buttons []uint, level int) int {
  // for _, b3 := range buttons {
  //   if b1^b2 == config {
  //     return level
  //   }
  // }
  return level
}

func parseMachines(strs []string) []Machine {
  p1 := regexp.MustCompile(`[\.|#]+`)
  p2 := regexp.MustCompile(`\(\S+\)+`)

  var machines []Machine
  for _, str := range strs {
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
    machines = append(machines, Machine{config, buttons})
  }

  return machines
}

