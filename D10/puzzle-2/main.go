package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	xRegister, cycleCounter := 1, 0

	for _, instruction := range instructionSet {
		if instruction.isAddx() {
			draw(xRegister, cycleCounter)
			cycleCounter++

			draw(xRegister, cycleCounter)
			cycleCounter++

			xRegister += instruction.value

			continue
		}

		draw(xRegister, cycleCounter)
		cycleCounter++
	}
}

func draw(registerX int, pixelPosition int) {
	pixelPosition = pixelPosition % 40 // Sprite can only move on the available 40 pixel width of the display
	spriteCenter := registerX

	if pixelPosition%40 == 0 {
		fmt.Println() // Display is 40 pixels wide, so go to the next CRT line here
	}

	if math.Abs(float64(spriteCenter)-float64(pixelPosition)) <= 1 {
		fmt.Print("#") // Lit pixel
	} else {
		fmt.Print(".") // Dark pixel
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
