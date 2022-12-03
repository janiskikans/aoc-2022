package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rucksacks := parseFileToSlice("../input.txt")

	itemSum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		commonItems := getCommonItems(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		itemSum += getItemPrioritySum(commonItems)
	}

	fmt.Println("Item sum", itemSum)
}

func getCommonItems(firstRucksack string, secondRucksack string, thirdRucksack string) []rune {
	var commonItems string

	for _, item := range firstRucksack {
		item := string(item)

		isAlreadyCommon := strings.Contains(commonItems, item)
		if strings.Contains(secondRucksack, item) && strings.Contains(thirdRucksack, item) && !isAlreadyCommon {
			commonItems += item
		}
	}

	return []rune(commonItems)
}

func getItemPrioritySum(items []rune) int {
	sum := 0

	for _, item := range items {
		// 97 and up all lower case letters in Unicode
		if item > 97 {
			sum += int(item) - 96

			continue
		}

		sum += int(item) - 38
	}

	return sum
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
