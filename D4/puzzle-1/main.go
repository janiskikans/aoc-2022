package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	pairAssignments := getInputAsSlice()

	overlappingPairCount := 0
	for _, pairAssignment := range pairAssignments {
		firstElfAssignment, secondElfAssignment := splitPairAssignment(pairAssignment)
		isFullOverlap := isFullOverlap(firstElfAssignment, secondElfAssignment)

		fmt.Println(firstElfAssignment, secondElfAssignment)
		fmt.Println("Is full overlap", isFullOverlap)

		if isFullOverlap {
			overlappingPairCount++
		}
	}

	fmt.Println("Fully overlapping pair count", overlappingPairCount)
}

func isFullOverlap(rangeOne []string, rangeTwo []string) bool {
	rangeOneStart, rangeOneEnd := getInt(rangeOne[0]), getInt(rangeOne[1])
	rangeTwoStart, rangeTwoEnd := getInt(rangeTwo[0]), getInt(rangeTwo[1])

	if (rangeOneStart >= rangeTwoStart && rangeOneEnd <= rangeTwoEnd) || (rangeTwoStart >= rangeOneStart && rangeTwoEnd <= rangeOneEnd) {
		return true
	}

	return false
}

func splitPairAssignment(pairAssignment string) ([]string, []string) {
	assignmentsByElf := strings.Split(pairAssignment, ",")

	var assignmentRanges [][]string
	for _, assignment := range assignmentsByElf {
		assignmentRanges = append(assignmentRanges, strings.Split(assignment, "-"))
	}

	return assignmentRanges[0], assignmentRanges[1]
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
