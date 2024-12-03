package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
  now := time.Now()
  defer func() {
    fmt.Printf("in %s\n", time.Now().Sub(now))
  }()

	fmt.Println("--- 3-2 ---")

  result := 0
  b, err := os.ReadFile("./input")
  if err != nil {
      fmt.Print(err)
  }

  str := string(b)
  r1 := regexp.MustCompile(`(?s)don't\(\)(.*?)+do\(\)`)
  matches := r1.FindAllStringSubmatch(str, -1)

  for _, match := range matches {
    str = strings.Replace(str, match[0], "", 1)
  }

  str = strings.SplitN(str, "don't", 2)[0]

  r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
  m := r.FindAllStringSubmatch(str, -1)

  for _, match := range m {
    v1, _ := strconv.Atoi(match[1])
    v2, _ := strconv.Atoi(match[2])

    fmt.Printf("mul(%d,%d)\n", v1, v2)
    result += (v1 * v2)
  }

	fmt.Printf("Result: %d\n", result)
}
