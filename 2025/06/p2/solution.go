package p2

import (
	"fmt"
  "strconv"
  "strings"

	"github.com/drcrees/aoc/helpers"
)

type Operation struct {
	operator string
	operands []string
}

func Solve() {
	fmt.Println("--- 6-2 ---")

	runes := helpers.ReadRunes("./2025/06/p2/input")
  ops := buildOperations(runes)

	result := 0
  for _, op := range ops {
    result += evaluate(op)
  }

	fmt.Printf("Result: %d\n", result)
}

func evaluate(op Operation) int {
  if op.operator == "*" {
    return op.multiply(0)
  }
  return op.add(0)
}

func (op *Operation) multiply(index int) int {
  if index == len(op.operands)-1 {
    a, _ := strconv.Atoi(op.operands[index])
    return a
  }
  if index+1 == len(op.operands)-1 {
    a, _ := strconv.Atoi(op.operands[index])
    b, _ := strconv.Atoi(op.operands[index+1])
    return a * b
  }
  a, _ := strconv.Atoi(op.operands[index])
  b, _ := strconv.Atoi(op.operands[index+1])
  return a * b * op.multiply(index+2)
}

func (op *Operation) add(index int) int {
  if index == len(op.operands)-1 {
    a, _ := strconv.Atoi(op.operands[index])
    return a
  }
  a, _ := strconv.Atoi(op.operands[index])
  return a + op.add(index+1)
}

func buildOperations(runes [][]rune) []Operation {
  var ops []Operation
  var operands []string
  var operator string

  operatorRow := len(runes)-1
  for col := len(runes[0])-1; col >= 0; col-- {
    operand := ""
    for row := 0; row < len(runes); row++ {
      if row == operatorRow && string(runes[row][col]) != " " {
        operator = string(runes[operatorRow][col])
        continue
      }
      operand += string(runes[row][col])
    }
    if strings.Trim(operand, " ") != "" {
      operands = append(operands, strings.Trim(operand, " "))
    }
    if col == 0 || strings.Trim(operand, " ") == "" {
      ops = append(ops, Operation{operator, operands})
      operands = []string{}
    }
  }
  return ops
}
