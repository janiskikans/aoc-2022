package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const NoopInstruction = "noop"
const AddxInstruction = "addx"

type Instruction struct {
	instruction string
	value       int
}

func main() {
	input := getInputAsSlice()
	instructionSet := parseInstructions(input)

	signalStrengthSum, xRegister, cycleCounter := 0, 1, 0

	for _, instruction := range instructionSet {
		if cycleCounter > 220 {
			break
		}

		if instruction.isAddx() {
			cycleCounter++
			incrementSingalStrength(&signalStrengthSum, xRegister, cycleCounter)

			cycleCounter++
			incrementSingalStrength(&signalStrengthSum, xRegister, cycleCounter)

			xRegister += instruction.value

			continue
		}

		cycleCounter++
		incrementSingalStrength(&signalStrengthSum, xRegister, cycleCounter)
	}

	fmt.Println("Signal strength sum", signalStrengthSum)
}

func incrementSingalStrength(strengthSum *int, xRegister int, cycleCounter int) {
	if (cycleCounter-20)%40 == 0 {
		*strengthSum += xRegister * cycleCounter
	}
}

func parseInstructions(input []string) []Instruction {
	var instructions []Instruction

	for _, line := range input {
		fields := strings.Fields(line)

		instructionValue := 0
		if len(fields) > 1 {
			instructionValue = getInt(fields[1])
		}

		instructions = append(instructions, Instruction{fields[0], instructionValue})
	}

	return instructions
}

func (instruction *Instruction) isAddx() bool {
	return instruction.instruction == AddxInstruction
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

func getInt(value string) int {
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return valueAsInt
}
