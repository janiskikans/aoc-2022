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

	commonItemPrioritySum := 0
	for _, rucksackContent := range rucksacks {
		compartmentOneContent, compartmentTwoContent := splitStringInTwo(rucksackContent)

		fmt.Println(compartmentOneContent, compartmentTwoContent)

		commonItems := getCommonItems(compartmentOneContent, compartmentTwoContent)
		commonItemPrioritySum += getItemPrioritySum(commonItems)

		fmt.Println(string(commonItems), getItemPrioritySum(commonItems))
	}

	fmt.Println("Common items priority sum", commonItemPrioritySum)
}

func splitStringInTwo(string string) (string, string) {
	middleIndex := len(string) / 2

	return string[0:middleIndex], string[middleIndex:]
}

// Runes? Really Go? Why...
func getCommonItems(compartmentOne string, compartmentTwo string) []rune {
	var commonItems string

	for _, item := range compartmentOne {
		item := string(item)

		isAlreadyCommon := strings.Contains(commonItems, item)
		if strings.Contains(compartmentTwo, item) && !isAlreadyCommon {
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
