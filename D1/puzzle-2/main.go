package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

		calorieSums[currentElfIndex] += calorieCount
	}

	// Sort calorie sums in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))

	topElfCalorieSum := 0
	for i := 0; i <= 2; i++ {
		topElfCalorieSum += calorieSums[i]
	}

	fmt.Println("Top elf calorie sum: ", topElfCalorieSum)
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
