package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const RockShape = "rock"
const PaperShape = "paper"
const ScissorsShape = "scissors"

var shapeMap = map[string]string{
	"A": RockShape,
	"X": RockShape,
	"B": PaperShape,
	"Y": PaperShape,
	"C": ScissorsShape,
	"Z": ScissorsShape,
}

var shapeScoreMap = map[string]int{
	RockShape:     1,
	PaperShape:    2,
	ScissorsShape: 3,
}

func main() {
	plays := parseFileToSlice("../input.txt")
	totalScore := 0

	for _, round := range plays {
		opponentShape, myShape := getRoundShapes(round)

		myShapeScore := shapeScoreMap[myShape]
		playScore := getPlayScore(myShape, opponentShape)

		totalScore += myShapeScore + playScore

		fmt.Println(myShape, opponentShape, myShapeScore, playScore, totalScore)
	}

	fmt.Println("Total score", totalScore)
}

func getRoundShapes(round string) (string, string) {
	roundShapes := strings.Fields(round)

	return shapeMap[roundShapes[0]], shapeMap[roundShapes[1]]
}

func getPlayScore(myShape string, opponentShape string) int {
	if myShape == opponentShape {
		return 3
	}

	if myShape == RockShape {
		if opponentShape == ScissorsShape {
			return 6
		}

		return 0
	}

	if myShape == PaperShape {
		if opponentShape == RockShape {
			return 6
		}

		return 0
	}

	if myShape == ScissorsShape && opponentShape == PaperShape {
		return 6
	}

	return 0
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
