package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	datastream := []rune(getInput())

	for i := 3; i < len(datastream); i++ {
		lastFourChars := []rune{datastream[i-3], datastream[i-2], datastream[i-1], datastream[i]}

		areDistinct := areCharsDistinct(lastFourChars)
		if areDistinct {
			fmt.Println("Message starts from:", i+1)

			break
		}
	}
}

func areCharsDistinct(chars []rune) bool {
	charMap := make(map[rune]bool)

	for _, char := range chars {
		exists := charMap[char]
		if exists {
			continue
		}

		charMap[char] = true
	}

	return len(charMap) == len(chars)
}

func getInput() string {
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

	return fileLines[0]
}
