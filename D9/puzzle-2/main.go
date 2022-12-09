package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	direction string
	moveCount int
}

type Coord struct {
	x int
	y int
}

const DirR = "R"
const DirL = "L"
const DirD = "D"
const DirU = "U"

func main() {
	moves := getMoves(getInputAsSlice())

	tailVisitedCoords := make(map[Coord]bool)

	// The rope now consists of 10 knots instead of just 2
	// Each counts as it's own Coord struct
	ropeKnots := make([]Coord, 10)

	tailStartCoord := ropeKnots[9]
	tailVisitedCoords[tailStartCoord] = true

	for _, move := range moves {
		for i := 0; i < move.moveCount; i++ {
			// Move head / 1st knot
			ropeKnots[0] = incrementCoords(ropeKnots[0], move.direction)

			// Go through all knots smilarly as the tail before
			for j := 0; j < len(ropeKnots)-1; j++ {
				// j+1 is the new tail in reference to j which is kinda the head for that knot pair
				ropeKnots[j+1] = getNewCoordsForTail(ropeKnots[j], ropeKnots[j+1])
			}

			tailVisitedCoords[ropeKnots[9]] = true
		}
	}

	fmt.Println("Visited by tail", len(tailVisitedCoords))
}

func incrementCoords(coord Coord, direction string) Coord {
	newCoord := Coord{x: coord.x, y: coord.y}

	if direction == DirD {
		newCoord.y--
	} else if direction == DirU {
		newCoord.y++
	} else if direction == DirL {
		newCoord.x--
	} else {
		newCoord.x++
	}

	return newCoord
}

func getNewCoordsForTail(head Coord, tail Coord) Coord {
	newTailCoord := Coord{tail.x, tail.y}

	xDiff := tail.x - head.x
	yDiff := tail.y - head.y

	if xDiff == 0 {
		if yDiff == -2 {
			newTailCoord.y++
		}

		if yDiff == 2 {
			newTailCoord.y--
		}

		return newTailCoord
	}

	if yDiff == 0 {
		if xDiff == -2 {
			newTailCoord.x++
		}

		if xDiff == 2 {
			newTailCoord.x--
		}

		return newTailCoord
	}

	if (xDiff < -1 || xDiff > 1) || (yDiff < -1 || yDiff > 1) {
		if xDiff > 0 {
			newTailCoord.x--
		}

		if xDiff <= 0 {
			newTailCoord.x++
		}

		if yDiff > 0 {
			newTailCoord.y--
		}

		if yDiff <= 0 {
			newTailCoord.y++
		}
	}

	return newTailCoord
}

func getMoves(input []string) []Move {
	var moves []Move
	for _, line := range input {
		splitLine := strings.Split(line, " ")

		moves = append(moves, Move{direction: splitLine[0], moveCount: getInt(splitLine[1])})
	}

	return moves
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
