package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const RockShape = "rock"
const PaperShape = "paper"
const ScissorsShape = "scissors"

const LoseResult = "X"
const DrawResult = "Y"
const WinResult = "Z"

var shapeMap = map[string]string{
	"A": RockShape,
	"B": PaperShape,
	"C": ScissorsShape,
}

var shapeScoreMap = map[string]int{
	RockShape:     1,
	PaperShape:    2,
	ScissorsShape: 3,
}

var resultScoreMap = map[string]int{
	LoseResult: 0,
	DrawResult: 3,
	WinResult:  6,
}

func main() {
	plays := parseFileToSlice("../input.txt")
	totalScore := 0

	for _, round := range plays {
		opponentShape, roundResult := getOpponentShapeAndResult(round)

		myShape, err := getMyShape(opponentShape, roundResult)
		if err != nil {
			fmt.Println(err)

			break
		}

		myShapeScore := shapeScoreMap[myShape]
		playScore := getPlayScore(myShape, opponentShape)

		totalScore += myShapeScore + playScore
	}

	fmt.Println("Total score", totalScore)
}

func getMyShape(opponentShape string, outcome string) (string, error) {
	outcomeScore := resultScoreMap[outcome]

	for shape := range shapeScoreMap {
		playScore := getPlayScore(shape, opponentShape)

		if playScore == outcomeScore {
			return shape, nil
		}
	}

	return "nil", errors.New("No shape calculated")
}

func getOpponentShapeAndResult(round string) (string, string) {
	data := strings.Fields(round)

	return shapeMap[data[0]], data[1]
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
