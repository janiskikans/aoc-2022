package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	treeWithHeighestScenicScore, score := getTreeWithHeighestScenicScore(trees)

	fmt.Println("Tree with heighest score", treeWithHeighestScenicScore, score)
}

func getTreeWithHeighestScenicScore(allTrees []Tree) (Tree, int) {
	var treeWithHeighestScore Tree

	heighestScore := 0
	for _, tree := range allTrees {
		scenicScore := getTreeScenicScore(tree, allTrees)

		if scenicScore > heighestScore {
			treeWithHeighestScore = tree
			heighestScore = scenicScore
		}
	}

	return treeWithHeighestScore, heighestScore
}

func getTreeScenicScore(currentTree Tree, allTrees []Tree) int {
	visibleTreesOnLeft := getVisibleTreesInDirection(currentTree, allTrees, DirLeft)
	visibleTreesOnRight := getVisibleTreesInDirection(currentTree, allTrees, DirRight)
	visibleTreesOnTop := getVisibleTreesInDirection(currentTree, allTrees, DirTop)
	visibleTreesOnBottom := getVisibleTreesInDirection(currentTree, allTrees, DirBottom)

	return len(visibleTreesOnBottom) * len(visibleTreesOnTop) * len(visibleTreesOnLeft) * len(visibleTreesOnRight)
}

func getVisibleTreesInDirection(currentTree Tree, allTrees []Tree, direction string) []Tree {
	allTreesInDirection := getTreesInDirection(currentTree, allTrees, direction)
	if len(allTreesInDirection) == 0 {
		return allTreesInDirection
	}

	var visibleTrees []Tree
	for i, tree := range allTreesInDirection {
		if tree.height <= currentTree.height {
			visibleTrees = append(visibleTrees, tree)

			if tree.height == currentTree.height {

				// We don't want to continue after first same height tree
				break
			}
		} else {
			if i < len(allTrees) {
				visibleTrees = append(visibleTrees, tree)

				break
			}
		}
	}

	return visibleTrees
}

func getTreesInDirection(currentTree Tree, allTrees []Tree, direction string) []Tree {
	var filteredTrees []Tree

	for _, tree := range allTrees {
		if tree.positionX == currentTree.positionX && tree.positionY == currentTree.positionY {
			continue
		}

		if direction == DirTop {
			if tree.positionY < currentTree.positionY && tree.positionX == currentTree.positionX {
				filteredTrees = append(filteredTrees, tree)
			}
		} else if direction == DirBottom {
			if tree.positionY > currentTree.positionY && tree.positionX == currentTree.positionX {
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

	// Sorting in the direction of how it's seen from the tree house
	sort.Slice(filteredTrees, func(a, b int) bool {
		if direction == DirLeft {
			return filteredTrees[a].positionX > filteredTrees[b].positionX
		} else if direction == DirRight {
			return filteredTrees[a].positionX < filteredTrees[b].positionX
		} else if direction == DirTop {
			return filteredTrees[a].positionY > filteredTrees[b].positionY
		} else {
			return filteredTrees[a].positionY < filteredTrees[b].positionY
		}
	})

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
