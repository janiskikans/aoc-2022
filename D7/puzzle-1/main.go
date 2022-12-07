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

	var currentPath []string
	directorySizes := make(map[string]int)
	numberRegex := getRegex("\\d+")

	for _, inputLine := range input {
		splitInput := strings.Split(inputLine, " ")

		// If we're on a "file line" we loop through each dir in the current path and add the file size
		// E.g. if the current path is ["/", "/a", "/a/e"] we're going to add the file size to all of those
		// because /a/e is child of /a and /
		if numberRegex.MatchString(inputLine) {
			for _, dir := range currentPath {
				if size, ok := directorySizes[dir]; ok {
					directorySizes[dir] = size + getInt(splitInput[0])

					continue
				}

				directorySizes[dir] = 0 + getInt(splitInput[0])
			}
		}

		// We remove the last dir of the current dir / "go up a dir"
		// Would be great if we would have native stacks in Go..
		if inputLine == "$ cd .." {
			currentPath = currentPath[0 : len(currentPath)-1]

			continue
		}

		if strings.HasPrefix(inputLine, "$ cd") {
			// Stupid way of correctly formatting directories, removing redundant slashes, etc.
			// Is ugly, but whatevs
			newPath := strings.Join(currentPath, "") + "/" + splitInput[2]
			currentPath = append(currentPath, strings.Replace(newPath, "//", "/", -1))

			continue
		}
	}

	fmt.Println("Directories")
	for dir, size := range directorySizes {
		fmt.Println(dir, "=>", size)
	}

	fmt.Println("\nDirectory (<= 100000) size sum:", getDirectorySizeSum(directorySizes, 100_000))
}

func getDirectorySizeSum(directories map[string]int, maxSizeLimit int) int {
	sum := 0

	for _, size := range directories {
		if size <= maxSizeLimit {
			sum += size
		}
	}

	return sum
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
