package main

import (
	"fmt"
	"os"
  "time"
  "regexp"
  "strconv"
)

func main() {
  now := time.Now()
  defer func() {
    fmt.Printf("in %s\n", time.Now().Sub(now))
  }()

	fmt.Println("--- 3-1 ---")

  result := 0
  b, err := os.ReadFile("./input")
  if err != nil {
      fmt.Print(err)
  }

  str := string(b)
  r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
  matches := r.FindAllStringSubmatch(str, -1)

  for _, match := range matches {
    v1, _ := strconv.Atoi(match[1])
    v2, _ := strconv.Atoi(match[2])

    result += (v1 * v2)
  }

	fmt.Printf("Result: %d\n", result)
}
