package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	height    int
	positionX int // Column
	positionY int // Row
}

const DirTop = "top"
const DirBottom = "bottom"
const DirLeft = "left"
const DirRight = "Right"

// WARNING! THIS IS QUITE UGLY. DON'T LOOK PLS

func main() {
	input := getInputAsSlice()

	var forestGrid [][]int
	for _, row := range input {
		var trees []int
		for _, tree := range strings.Split(row, "") {
			trees = append(trees, getInt(tree))
		}

		forestGrid = append(forestGrid, trees)
	}

	columnCount := len(forestGrid[0])
	rowCount := len(forestGrid)

	var trees []Tree

	// Making Tree structures
	for row := 0; row < rowCount; row++ {
		for col := 0; col < columnCount; col++ {
			tree := Tree{height: forestGrid[row][col], positionX: col, positionY: row}

			trees = append(trees, tree)
		}
	}

	countOfVisibleTrees := 0
	for _, tree := range trees {
		if isTallestTreeInAnyDirection(tree, trees) {
			countOfVisibleTrees++
		}
	}

	fmt.Println("Count of visible trees", countOfVisibleTrees)
}

func isTallestTreeInAnyDirection(currentTree Tree, allTrees []Tree) bool {
	tallerTreesOnLeft := getTallerOrSameHeightTreesThanHeight(
		currentTree.height,
		getTreesInDirection(currentTree, allTrees, DirLeft),
	)

	if len(tallerTreesOnLeft) == 0 {
		return true
	}

	tallerTreesOnRight := getTallerOrSameHeightTreesThanHeight(
		currentTree.height,
		getTreesInDirection(currentTree, allTrees, DirRight),
	)

	if len(tallerTreesOnRight) == 0 {
		return true
	}

	tallerTreesOnTop := getTallerOrSameHeightTreesThanHeight(
		currentTree.height,
		getTreesInDirection(currentTree, allTrees, DirTop),
	)

	if len(tallerTreesOnTop) == 0 {
		return true
	}

	tallerTreesOnBottom := getTallerOrSameHeightTreesThanHeight(
		currentTree.height,
		getTreesInDirection(currentTree, allTrees, DirBottom),
	)

	return len(tallerTreesOnBottom) == 0
}

func getTreesInDirection(currentTree Tree, allTrees []Tree, direction string) []Tree {
	var filteredTrees []Tree

	for _, tree := range allTrees {
		if direction == DirTop {
			if tree.positionY > currentTree.positionY && tree.positionX == currentTree.positionX {
				filteredTrees = append(filteredTrees, tree)
			}
		} else if direction == DirBottom {
			if tree.positionY < currentTree.positionY && tree.positionX == currentTree.positionX {
				filteredTrees = append(filteredTrees, tree)
			}
		} else if direction == DirLeft {
			if tree.positionX < currentTree.positionX && tree.positionY == currentTree.positionY {
				filteredTrees = append(filteredTrees, tree)
			}
		} else if tree.positionX > currentTree.positionX && tree.positionY == currentTree.positionY {
			filteredTrees = append(filteredTrees, tree)
		}
	}

	return filteredTrees
}

func getTallerOrSameHeightTreesThanHeight(height int, trees []Tree) []Tree {
	var filteredTrees []Tree

	for _, tree := range trees {
		if tree.height >= height {
			filteredTrees = append(filteredTrees, tree)
		}
	}

	return filteredTrees
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
