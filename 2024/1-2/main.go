package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strconv"
  "sort"
  "strings"
  "time"
)

func main() {
  now := time.Now()
  defer func() {
    fmt.Printf("in %s\n", time.Now().Sub(now))
  }()

	fmt.Println("--- 1-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int = 0
  var l1, l2 []int
  var m map[int]int = make(map[int]int)
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
    n := l1[i]

    count := 0
    for j := 0; j < len(l2); j++ {
      if n == l2[j] {
        count++
      }
    }
    m[n] = count
  }

  for i, v := range(m) {
    result += i * v
  }

	fmt.Printf("Result: %d\n", result)
}

func parse(str string) (int, int) {
  numStrs := strings.Split(str, "   ")

  n1, _ := strconv.Atoi(numStrs[0])
  n2, _ := strconv.Atoi(numStrs[1])

  return n1, n2;
}

