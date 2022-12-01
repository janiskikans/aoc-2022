package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := parseFileToSlice("../input.txt")

	calorieSums := []int{0}
	currentElfIndex := 0

	for _, line := range lines {
		calorieCount, err := strconv.Atoi(line)
		if err != nil {
			currentElfIndex++
			calorieSums = append(calorieSums, 0)

			continue
		}

		calorieSums[currentElfIndex] = calorieSums[currentElfIndex] + calorieCount
	}

	// Get and output max calorie count
	maxCalorieSum, elemIndex := getMaxFromSlice(calorieSums)

	fmt.Println("Elf index: ", elemIndex, ", Calorie sum: ", maxCalorieSum)
}

func parseFileToSlice(path string) []string {
	file, err := os.Open(path)
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

func getMaxFromSlice(sliceOfNumbers []int) (int, int) {
	max := 0
	maxElemIndex := 0

	for index, number := range sliceOfNumbers {
		if number > max {
			max = number
			maxElemIndex = index
		}
	}

	return max, maxElemIndex
}
