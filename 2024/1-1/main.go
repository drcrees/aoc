package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strconv"
  "sort"
  "strings"
  "math"
  "time"
)

func main() {
  now := time.Now()
  defer func() {
    fmt.Printf("in %s\n", time.Now().Sub(now))
  }()

	fmt.Println("--- 1-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int = 0
  var l1, l2 []int
	for scanner.Scan() {
    n1, n2 := parse(scanner.Text())

    l1 = append(l1, n1)
    l2 = append(l2, n2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

  sort.Ints(l1)
  sort.Ints(l2)

  for i := 0; i < len(l1); i++ {
    d := math.Abs(float64(l1[i] - l2[i]))
    result += int(d)
  }

	fmt.Printf("Result: %d\n", result)
}

func parse(str string) (int, int) {
  numStrs := strings.Split(str, "   ")

  n1, _ := strconv.Atoi(numStrs[0])
  n2, _ := strconv.Atoi(numStrs[1])

  return n1, n2;
}

