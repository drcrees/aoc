package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strconv"
  "strings"
  "time"
)

func main() {
  now := time.Now()
  defer func() {
    fmt.Printf("in %s\n", time.Now().Sub(now))
  }()

	fmt.Println("--- 2-1 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int = 0
	for scanner.Scan() {
    if parse(scanner.Text()) {
      result++
    }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", result)
}

func parse(str string) bool {
  numStrs := strings.Split(str, " ")

  return IsSafe(numStrs, 0, 0, false)
}

func IsSafe(numStrs []string, index int, dir int, isSafe bool) bool {
  if index == len(numStrs) - 1 {
    return isSafe
  }

  n1, _ := strconv.Atoi(numStrs[index])
  n2, _ := strconv.Atoi(numStrs[index + 1]) 

  diff := n1 - n2 
  if diff == 0 || diff < -3 || diff > 3 {
    return false
  }

  // change in direction
  if (dir < 0 && diff > 0) || (dir > 0 && diff < 0){
    return false 
  }

  return IsSafe(numStrs, index+1, diff, true)
}
