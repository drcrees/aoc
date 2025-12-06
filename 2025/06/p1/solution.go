package p1

import (
	"fmt"
  "regexp"
  "strconv"

	"github.com/drcrees/aoc/helpers"
)

type Operation struct {
	operator string
	operands []int
}

func Solve() {
	fmt.Println("--- 6-1 ---")

	strs := helpers.ReadLines("./2025/06/p1/input")
  ops := buildOperations(strs)

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
    return op.operands[index]
  }
  if index+1 == len(op.operands)-1 {
    return op.operands[index] * op.operands[index+1]
  }
  return op.operands[index] * op.operands[index+1] * op.multiply(index+2)
}

func (op *Operation) add(index int) int {
  if index == len(op.operands)-1 {
    return op.operands[index]
  }
  return op.operands[index] + op.add(index+1)
}

func buildOperations(strs []string) []Operation {
  p1 := regexp.MustCompile(`\d+`)
  p2 := regexp.MustCompile(`\S`)

  var operands [][]string
  var operators []string

  for _, str := range strs {
    operands = append(operands, p1.FindAllString(str, -1))
    operators = p2.FindAllString(str, -1)
  }

  var ops []Operation
  for i, operator := range operators {
    operation := Operation{operator: operator}
    for _, operands := range operands {
      if len(operands) > 0 {
        operand, _ := strconv.Atoi(operands[i])
        operation.operands = append(operation.operands, operand)
      }
    }
    ops = append(ops, operation)
  }

  return ops
}
