package main

import (
	"fmt"
	"sort"
	"strconv"
)

const OPERATION_MULTIPLY = "multiply"
const OPERATION_PLUS = "plus"

const OPERATION_ARG_OLD = "old"

type Operation struct {
	operator string
	arg1     string
	arg2     string
}

type Test struct {
	divisionOperand int
	truthyTargetId  int
	falsyTargetId   int
}

type Monkey struct {
	id                 int
	items              []int
	operation          Operation
	test               Test
	inspectedItemCount int
}

func main() {
	monkeys := buildMonkeys()

	for round := 1; round <= 20; round++ {
		fmt.Println("______ ROUND", round, "_______")

		for monkeyIndex := range monkeys {
			monkey := &monkeys[monkeyIndex]

			for _, item := range monkey.items {
				monkey.inspectedItemCount++

				newWorryLevel := getNewWorryLevel(item, monkey.operation)
				item = newWorryLevel / 3

				isDivisible := isDivisibleBy(item, monkey.test.divisionOperand)

				if isDivisible {
					toMonkey := &monkeys[monkey.test.truthyTargetId]
					toMonkey.items = append(toMonkey.items, item)

					continue
				}

				toMonkey := &monkeys[monkey.test.falsyTargetId]
				toMonkey.items = append(toMonkey.items, item)
			}

			monkey.items = []int{}
		}

		for _, monkey := range monkeys {
			fmt.Println("Monkey", monkey.id, ":", monkey.items, "Inspected item count:", monkey.inspectedItemCount)
		}
	}

	sort.Slice(monkeys, func(a, b int) bool {
		return monkeys[a].inspectedItemCount > monkeys[b].inspectedItemCount
	})

	monkeyBusinessLevel := monkeys[0].inspectedItemCount * monkeys[1].inspectedItemCount
	fmt.Println("\nMonkey business level", monkeyBusinessLevel)
}

func isDivisibleBy(dividend int, divisor int) bool {
	return dividend%divisor == 0
}

func getNewWorryLevel(item int, operation Operation) int {
	firstOperand, secondOperand := getOperands(item, operation)

	if operation.operator == OPERATION_MULTIPLY {
		return firstOperand * secondOperand
	}

	return firstOperand + secondOperand
}

func getOperands(item int, operation Operation) (int, int) {
	firstOperand := getOperand(item, operation.arg1)
	secondOperand := getOperand(item, operation.arg2)

	return firstOperand, secondOperand
}

func getOperand(item int, operand string) int {
	if operand == OPERATION_ARG_OLD {
		return item
	}

	return getInt(operand)
}

func buildMonkeys() []Monkey {
	return []Monkey{
		{0, []int{83, 88, 96, 79, 86, 88, 70}, Operation{OPERATION_MULTIPLY, OPERATION_ARG_OLD, "5"}, Test{11, 2, 3}, 0},
		{1, []int{59, 63, 98, 85, 68, 72}, Operation{OPERATION_MULTIPLY, OPERATION_ARG_OLD, "11"}, Test{5, 4, 0}, 0},
		{2, []int{90, 79, 97, 52, 90, 94, 71, 70}, Operation{OPERATION_PLUS, OPERATION_ARG_OLD, "2"}, Test{19, 5, 6}, 0},
		{3, []int{97, 55, 62}, Operation{OPERATION_PLUS, OPERATION_ARG_OLD, "5"}, Test{13, 2, 6}, 0},
		{4, []int{74, 54, 94, 76}, Operation{OPERATION_MULTIPLY, OPERATION_ARG_OLD, OPERATION_ARG_OLD}, Test{7, 0, 3}, 0},
		{5, []int{58}, Operation{OPERATION_PLUS, OPERATION_ARG_OLD, "4"}, Test{17, 7, 1}, 0},
		{6, []int{66, 63}, Operation{OPERATION_PLUS, OPERATION_ARG_OLD, "6"}, Test{2, 7, 5}, 0},
		{7, []int{56, 56, 90, 96, 68}, Operation{OPERATION_PLUS, OPERATION_ARG_OLD, "7"}, Test{3, 4, 1}, 0},
	}
}

func getInt(value string) int {
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return valueAsInt
}
