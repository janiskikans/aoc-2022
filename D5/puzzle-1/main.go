package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := getInputAsSlice()

	// Creating stacks and filling them up
	stackCount := getStackCount(input)
	stacks := make([]Stack, stackCount)

	instructionStartRow := 0
	numberAndSpaceRegex := getRegex("^[0-9 ]+$")
	for i, inputLine := range input {
		if inputLine == "" || numberAndSpaceRegex.MatchString(inputLine) {
			// We have reached end of stack defs
			instructionStartRow = i + 3

			break
		}

		currentStack := 0
		for i := 1; i < stackCount*4-2; i += 4 {
			crateLetter := inputLine[i]
			if crateLetter == 32 {
				currentStack++

				continue
			}

			stacks[currentStack].addCrateToBottom(rune(crateLetter))
			currentStack++
		}
	}

	for i, instructionLine := range input {
		if i < instructionStartRow-1 {
			continue
		}

		movableCrateCount, fromCrateStackIndex, toCreateStackIndex := getInstructionValues(instructionLine)
		moveCrates(movableCrateCount, &stacks[fromCrateStackIndex-1], &stacks[toCreateStackIndex-1])

		outputAllStacks(stacks)
	}
}

func moveCrates(crateCount int, fromStack *Stack, toStack *Stack) {
	for i := 0; i < crateCount; i++ {
		toStack.pushCrate(fromStack.popCrate())
	}
}

func outputAllStacks(stacks []Stack) {
	for i, stack := range stacks {
		fmt.Print(i+1, " ")
		stack.output()
	}

	fmt.Println("-------")
}

func getInstructionValues(instructionLine string) (int, int, int) {
	instructionNumbersAsStrings := getRegex("[0-9]+").FindAllString(instructionLine, -1)

	var instructionNumbers []int
	for _, numberAsString := range instructionNumbersAsStrings {
		instructionNumbers = append(instructionNumbers, getInt(numberAsString))
	}

	return instructionNumbers[0], instructionNumbers[1], instructionNumbers[2]
}

func getStackCount(input []string) int {
	numberRegex := getRegex("^[0-9 ]+$")

	for _, inputLine := range input {
		if !numberRegex.MatchString(inputLine) {
			continue
		}

		stackNumbers := strings.Fields(inputLine)
		maxStackNumber := 0
		for _, stackNumber := range stackNumbers {
			stackNumber := getInt(stackNumber)

			if stackNumber > maxStackNumber {
				maxStackNumber = stackNumber
			}
		}

		return maxStackNumber
	}

	return 0
}

func getRegex(regexExpreession string) *regexp.Regexp {
	return regexp.MustCompile(regexExpreession)
}

func getInt(value string) int {
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return valueAsInt
}

func getInputAsSlice() []string {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	file.Close()

	return fileLines
}

// Stack struct
type Stack struct {
	crates []rune
}

func (stack *Stack) addCrateToBottom(crate rune) {
	stack.crates = append([]rune{crate}, stack.crates...)
}

func (stack *Stack) popCrate() rune {
	length := len(stack.crates)
	if length == 0 {
		panic("No crates left in stack")
	}

	crateToRemove := stack.crates[length-1]
	stack.crates = stack.crates[0 : length-1]

	return crateToRemove
}

func (stack *Stack) pushCrate(crate rune) {
	stack.crates = append(stack.crates, crate)
}

func (stack Stack) output() {
	for _, crate := range stack.crates {
		fmt.Print("[", string(crate), "] ")
	}

	fmt.Println()
}
