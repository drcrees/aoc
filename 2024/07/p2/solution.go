package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/drcrees/aoc/helpers"
)

type Equation struct {
	answer   int64
	operands []int64
}

func Solve() {
	fmt.Println("--- 7-2 ---")
	strs := helpers.ReadDelimitedStrings("./2024/07/p2/input", ": ")
	equations := parse(strs)

	var result int64
	for _, equation := range equations {
		if equation.isValid(0, 0) {
			result += equation.answer
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func (eq *Equation) isValid(result int64, index int) bool {
	if index == len(eq.operands) {
		return result == eq.answer
	}

	if result > eq.answer {
		return false
	}

	p1 := int64(result) + int64(eq.operands[index])
	p2 := int64(result) * int64(eq.operands[index])
	p3, _ := strconv.ParseInt(strconv.FormatInt(result, 10)+strconv.FormatInt(eq.operands[index], 10), 10, 64)

	return eq.isValid(p1, index+1) || eq.isValid(p2, index+1) || eq.isValid(p3, index+1)
}

func parse(strs [][]string) (e []*Equation) {
	for _, v := range strs {
		answer, _ := strconv.ParseInt(v[0], 10, 64)
		var operands []int64
		operandStrs := strings.Split(v[1], " ")
		for _, o := range operandStrs {
			operand, _ := strconv.ParseInt(o, 10, 64)
			operands = append(operands, operand)
		}

		e = append(e, &Equation{
			answer:   answer,
			operands: operands,
		})
	}

	return e
}
